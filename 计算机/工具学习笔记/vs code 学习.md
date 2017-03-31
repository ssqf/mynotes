# Visual Studio Code 学习
---
## 常用命令
`Ctrl + Shift +j` 高级搜索  
`Ctrl + p` 打开文件搜索  
`Ctrl + Shift + o` 打开符号跳转 只在当前文件
`Ctrl + Shift + p` 当前编辑器命令
`Ctrl + G` 行跳转  
`Ctrl + Tab` 打开文件切换 从先到后  
`Ctrl + Shift + tab` 打开文件切换 从后到先  
`Alt + Left` 和  `Alt+Right` 在打开的文件中左右切换  
`Ctrl + Click` 跳转到定义  和 `Ctrl + Alt + Click` 新窗口打开定义  `Ctrl + K` 和 `Ctrl + Left` 前后转  
`Ctrl + T` 通过名字打开符号 在当前工作空间搜索  
`Ctrl + Shift + \` 括号匹配  
`Ctrl + Click ` 列模式选择  `Ctrl+Alt+Down` or `Ctrl+Alt+Up` 列选择   
`Ctrl+Shift+L` `Ctrl+F2`  选择一个词  
`Ctrl+Shift+[` 和 `Ctrl+Shift+]` 折叠和展开代码块        
`Ctrl+K` `Ctrl+0` 折叠全部(按住Ctrl 再分别按k 和 0 算一条命令)  和 `Ctrl+K` `Ctrl+J` 展开全部 `Ctrl+K` `Ctrl+2` 展开2级折叠       
`Ctrl+` \`   打开集成终端  `Ctrl+Ins` and `Shift+Ins` 和 `Ctrl+Shift+C` and `Ctrl+Shift+V` 在终端复制粘贴     
`Ctrl+-` 和  `Ctrl+=` 缩小 放大      
`Alt+F12`  临时预览函数定义

## 常用设置
1、修改vscode的markdown-it 打开extension.js 中 breaks 选型 `html: true, breaks: true,`
    C:\Program Files (x86)\Microsoft VS Code\resources\app\extensions\markdown\out\markdownEngine.js  

2、使用vscode-paste-image 插件时遇到的问题，和修改方法
一、windows 下产生 .\image\xxx.png 预览的时候不能显示图片

需要变为 ./image/xxx.png 才可正常

	/**
     * Convert the given Windows or Unix-style path into a normalized path that only uses forward slashes and has all superflous '..' sequences removed.
     * If the path starts with a Windows-style drive letter, a '/' is prepended.
     */
    Paster.normalize = function(pathStr) {
        pathStr = pathStr.replace(/\\/g, '/');
		if (/^[a-zA-Z]\:\//.test(pathStr)) {
            pathStr = '/' + pathStr;
        } 
        pathStr = path.normalize(pathStr); // use node's normalize to remove '<dir>/..' etc.
        pathStr = pathStr.replace(/\\/g, '/');
        return pathStr;
    };
二、相对目录使用相对于工程根目录的相对路径较好

绝对路径就使用绝对路径
'.' 开头的使用文件相对路径
非点开头的非绝对路径使用相对工程根目录的相对路径
        // generate image path
        if (path.isAbsolute(folderPathFromConfig)) {
            imagePath = path.join(folderPathFromConfig, imageFileName);
        }
        else if('.' === folderPathFromConfig.charAt(0)){
            imagePath = path.join(folderPath, folderPathFromConfig, imageFileName);
        }
        else {
            imagePath = path.join(vscode.workspace.rootPath, folderPathFromConfig, imageFileName);
        }
        return imagePath;
```js
'use strict';
var vscode = require('vscode');
var path = require('path');
var fs = require('fs');
var child_process_1 = require('child_process');
var moment = require('moment');
function activate(context) {
    console.log('Congratulations, your extension "vscode-paste-image" is now active!');
    var disposable = vscode.commands.registerCommand('extension.pasteImage', function () {
        Paster.paste();
    });
    context.subscriptions.push(disposable);
}
exports.activate = activate;
function deactivate() {
}
exports.deactivate = deactivate;
var Paster = (function () {
    function Paster() {
    }
    Paster.paste = function () {
        var _this = this;
        // get current edit file path
        var editor = vscode.window.activeTextEditor;
        if (!editor)
            return;
        var fileUri = editor.document.uri;
        if (!fileUri)
            return;
        if (fileUri.scheme === 'untitled') {
            vscode.window.showInformationMessage('Before paste image, you need to save current edit file first.');
            return;
        }
        // get selection as image file name, need check
        var selection = editor.selection;
        var selectText = editor.document.getText(selection);
        if (selectText && !/^[\w\-.]+$/.test(selectText)) {
            vscode.window.showInformationMessage('Your selection is not a valid file name!');
            return;
        }
        // get image destination path
        var folderPathFromConfig = vscode.workspace.getConfiguration('pasteImage')['path'];
        if (folderPathFromConfig && (folderPathFromConfig.length !== folderPathFromConfig.trim().length)) {
            vscode.window.showErrorMessage('The specified path is invalid. "' + folderPathFromConfig + '"');
            return;
        }
        var filePath = fileUri.fsPath;
        var imagePath = this.getImagePath(filePath, selectText, folderPathFromConfig);
        //vscode.window.showErrorMessage("imagePath:" + imagePath);
        this.createImageDirWithImagePath(imagePath).then(function (imagePath) {
            // save image and insert to current edit file
            //vscode.window.showErrorMessage("imagePath:" + imagePath);
            _this.saveClipboardImageToFileAndGetPath(imagePath, function (imagePath) {
                if (!imagePath)
                    return;
                if (imagePath === 'no image') {
                    return;
                }
                imagePath = _this.renderFilePath(editor.document.languageId, filePath, imagePath);
                editor.edit(function (edit) {
                    var current = editor.selection;
                    if (current.isEmpty) {
                        edit.insert(current.start, imagePath);
                    }
                    else {
                        edit.replace(current, imagePath);
                    }
                });
            });
        }).catch(function (err) {
            vscode.window.showErrorMessage('Failed make folder.');
            return;
        });
    };
    Paster.getImagePath = function (filePath, selectText, folderPathFromConfig) {
        // image file name
        var imageFileName = "";
        if (!selectText) {
            imageFileName = moment().format("Y-MM-DD-HH-mm-ss") + ".png";
        }
        else {
            imageFileName = selectText + ".png";
        }
        // image output path
        var folderPath = path.dirname(filePath);
        var imagePath = "";
        // generate image path
        if (path.isAbsolute(folderPathFromConfig)) {
            imagePath = path.join(folderPathFromConfig, imageFileName);
        }
        else if('.' === folderPathFromConfig.charAt(0)){
            imagePath = path.join(folderPath, folderPathFromConfig, imageFileName);
        }
        else {
            imagePath = path.join(vscode.workspace.rootPath, folderPathFromConfig, imageFileName);
        }
        return imagePath;
    };
    /**
     * create directory for image when directory does not exist
     */
    Paster.createImageDirWithImagePath = function (imagePath) {
        return new Promise(function (resolve, reject) {
            var imageDir = path.dirname(imagePath);
            fs.exists(imageDir, function (exists) {
                if (exists) {
                    resolve(imagePath);
                    return;
                }
                fs.mkdir(imageDir, function (err) {
                    if (err) {
                        reject(err);
                        return;
                    }
                    resolve(imagePath);
                });
            });
        });
    };
    /**
     * use applescript to save image from clipboard and get file path
     */
    Paster.saveClipboardImageToFileAndGetPath = function (imagePath, cb) {
        if (!imagePath)
            return;
        var platform = process.platform;
        if (platform === 'win32') {
            // Windows
            var scriptPath = path.join(__dirname, '../../res/pc.ps1');
            var powershell = child_process_1.spawn('powershell', [
                '-noprofile',
                '-noninteractive',
                '-nologo',
                '-sta',
                '-executionpolicy', 'unrestricted',
                '-windowstyle', 'hidden',
                '-file', scriptPath,
                imagePath
            ]);
            powershell.on('exit', function (code, signal) {
                // console.log('exit', code, signal);
            });
            powershell.stdout.on('data', function (data) {
                cb(data.toString().trim());
            });
        }
        else if (platform === 'darwin') {
            // Mac
            var scriptPath = path.join(__dirname, '../../res/mac.applescript');
            var ascript = child_process_1.spawn('osascript', [scriptPath, imagePath]);
            ascript.on('exit', function (code, signal) {
                // console.log('exit',code,signal);
            });
            ascript.stdout.on('data', function (data) {
                cb(data.toString().trim());
            });
        }
        else {
            // Linux 
            var scriptPath = path.join(__dirname, '../../res/linux.sh');
            var ascript = child_process_1.spawn('sh', [scriptPath, imagePath]);
            ascript.on('exit', function (code, signal) {
                // console.log('exit',code,signal);
            });
            ascript.stdout.on('data', function (data) {
                var result = data.toString().trim();
                if (result == "no xclip") {
                    vscode.window.showInformationMessage('You need to install xclip command first.');
                    return;
                }
                cb(result);
            });
        }
    };

	/**
     * Convert the given Windows or Unix-style path into a normalized path that only uses forward slashes and has all superflous '..' sequences removed.
     * If the path starts with a Windows-style drive letter, a '/' is prepended.
     */
    Paster.normalize = function(pathStr) {
        pathStr = pathStr.replace(/\\/g, '/');
		if (/^[a-zA-Z]\:\//.test(pathStr)) {
            pathStr = '/' + pathStr;
        } 
        pathStr = path.normalize(pathStr); // use node's normalize to remove '<dir>/..' etc.
        pathStr = pathStr.replace(/\\/g, '/');
        if('.' != pathStr.charAt(0) && '/' != pathStr.charAt(0))
        {
            pathStr = "./" + pathStr;
        }
        return pathStr;
    };
    /**
     * render the image file path dependen on file type
     * e.g. in markdown image file path will render to ![](path)
     */
    Paster.renderFilePath = function (languageId, docPath, imageFilePath) {
        imageFilePath = path.relative(path.dirname(docPath), imageFilePath);
        if (languageId === 'markdown') {
            return "![](" + this.normalize(imageFilePath) + ")";
        }
        else {
            return imageFilePath;
        }
    };
    return Paster;
}());
//# sourceMappingURL=extension.js.map
```

```powershell
param($imagePath)

# Adapted from https://github.com/octan3/img-clipboard-dump/blob/master/dump-clipboard-png.ps1

Add-Type -Assembly PresentationCore
$img = [Windows.Clipboard]::GetImage()

if ($img -eq $null) {
    "no image"
    Exit 1
}

if (-not $imagePath) {
    "no image path"
    Exit 1
}

$fcb = new-object Windows.Media.Imaging.FormatConvertedBitmap($img, [Windows.Media.PixelFormats]::Rgb24, $null, 0)
$stream = [IO.File]::Open($imagePath, "OpenOrCreate")
$encoder = New-Object Windows.Media.Imaging.PngBitmapEncoder
$encoder.Frames.Add([Windows.Media.Imaging.BitmapFrame]::Create($fcb)) | out-null
$encoder.Save($stream) | out-null
$stream.Dispose() | out-null

chcp 65001 | out-null
$imagePath
```