# golang工程管理工具

## go命令
1. go常用子命令
```
	build            compile packages and dependencies
	clean            remove object files
	doc              show documentation for package or symbol
	env              print Go environment information
	fmt              run gofmt on package sources
	get              download and install packages and dependencies
	install          compile and install packages and dependencies
	list             list packages
	run              compile and run Go program
	test             test packages
	version          print Go version
	vet              run go tool vet on packages

Use "go help [command]" for more information about a command.
```

## gofmt命令
1. 使用`gofmt [flags] [path ...]`path可以是文件和目录，目录则处理所有go源文件
    - `-d`比较原始和新调整后的差异，需要有diff工具可用 
    - `-e`打印所有错误
    - `-l`不输出重新格式化的源码，输出格式化后不同的文件名
    - `-w`不输出重新格式化的源码,并覆盖原文件
    - `-r`按规则替换
    - `-s`尝试简化代码

## godoc命令
1. 启动本地文档服务`godoc -http=:6060`
2. 提取代码文档`godoc package [name ...]`

## vscode golang插件
- 代码自动格式化，可以选择gofmt和goreturns
- 自动添加移除包依赖
- 代码调试

## 交叉编译
设置相关环节变量即可
GOOS：目标可执行程序运行操作系统，支持 darwin，freebsd，linux，windows  
GOARCH：目标可执行程序操作系统构架，包括 386，amd64，arm  
Golang version 1.5以前版本在首次交叉编译时还需要配置交叉编译环境：  
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ./make.bash  
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 ./make.bash  

### powershell 设置变量
1. linxu 64位系统程序
    $env:GOOS="linux"
	$env:GOARCH="amd64"
2. windows 64位程序
	$env:GOOS="windows"  
	$env:GOARCH="amd64"  
