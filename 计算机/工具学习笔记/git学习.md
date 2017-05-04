# git学习笔记

[TOC]

## 〇、基础介绍
* **分布式版本管理**
> 在本地有完整的服务器端文件克隆，包括文件及每一次的版本信息
> 可在本地修改本提交

* **基本概念**  
git内文件有*三种状态*：已提交（committed），已修改（modified）和已暂存（staged）


## 一、常用命令

Git 提供了一个叫做 git config 的工具（译注：实际是 git-config 命令，只不过可以通过 git 加一个名字来呼叫此命令。），专门用来配置或读取相应的工作环境变量。而正是由这些环境变量，决定了 Git 在各个环节的具体工作方式和行为。

### 1.1 用户信息

git config --global user.name "John Doe"   
git config --global user.email johndoe@example.com   
用于配置个人信息，每次提交都附带

### 1.2 查看配置信息

git config --list    查看各种配置信息

### 1.3 获取帮助

git help \<verb>   
git \<verb> --help   
man git-\<verb>  
如：`git help config` 获取config命令的使用

### 1.4 获取项目仓库 git init / git clone
* **从本地初始化**   
    git init  初始化本地项目生成相应.git配置管理目录   
    **文件纳入版本控制**  
    `git add *.c` 添加所有C文件  
    `git add README`  添加README  
    `git add .`  添加所有文件   
    `git commit -m 'initial project version`  提交到版本管理库中
* **从现有仓库克隆**   
    `git clone git://github.com/schacon/grit.git` 把URL的开源项目克隆的本地生成grit录   
    `git clone git://github.com/schacon/grit.git mygrit` 指定生成的目录`mygrit` 要是空目录  

### 1.5 记录更新到仓库 git status
>工作目录下文件只有两种状态：已跟踪或未跟踪。暂存区:修改后还没有提交
    
`git status`检查当前状态,`On branch master`表示在master分支上，`Untracked files:`表示文件未被加入跟踪，`Changes to be committed:`加入跟踪收暂存区但未提交，`Changed but not updated:`已经暂存但未提交  
`git add README`新文件README加入跟踪，git不会自动把新文件加入库，还可以把修改的文件收入暂存 `git add .`可以添加所有文件     

### 1.6 忽略文件 .gitignore 
在工作目录创建`.gitignore`文件。在其中加入要忽略的文件，相关文将不会再出现在未跟踪文件列表中。每行一类文件。

```shell
# 此为注释 – 将被 Git 忽略
*.a # 忽略所有 .a 结尾的文件
!lib.a # 但 lib.a 除外
/TODO # 仅仅忽略项目根目录下的 TODO 文件，不包括 subdir/TODO
build/ # 忽略 build/ 目录下的所有文件
doc/*.txt # 会忽略 doc/notes.txt 但不包括 doc/server/arch.txt
*.[oa] #.a或.o结尾的文件
```

glob 模式是指 shell 所使用的简化了的正则表达式。星号（ *）匹配零个或多个任意字符； [abc] 匹配任何一个列在方括号中的字符（这个例子要么匹配一个 a，要么匹配一个 b，要么匹配一个 c）；问号（ ?）只匹配一个任意字符；如果在方括号中使用短划线分隔两个字符，表示所有在这两个字符范围内的都可以匹配（比如 [0-9] 表示匹配所有 0 到9 的数字）

如果文件已经添加到追踪里，需要先删除 `git rm --cached filename`

### 1.7 查看已暂存和未暂存的更新 git diff

`git diff`比较的是工作目录中当前文件和暂存区域快照之间的差异，也就是修改之后还没有暂存起来的变化内容 *会查看完整内容q可以退出*   
`git diff --cached` 已经暂存起来的文件和上次提交时的快照之间的差异 `git diff --staged` 1.6版本之后的git可以上使用效果一样

### 1.8 提交更新 git commit

`git commit` 启动默认编辑器输入提交信息 `git config --global core.editor` 设定编辑器   
`git commit -m "这里是提交的说明"` -m 参数后跟提交说明的方式，在一行命令中提交更新,提交的是暂存区的快照，没有添加到暂存区的是不会被提交，而是继续保持修改状态  
`git commit -a ` 加上 -a 选项，Git 就会自动把所有已经跟踪过的文件暂存起来一并提交，从而跳过 git add 步骤

### 1.8 删除文件 git rm

如果只是简单的从工作目录删除，则会出现`Changed but not updated` 不会从版本跟踪中删除   
`git rm file` 把file从版本跟踪中删除，如果删除前已经修改并加入了暂存区，则可以使用  `-f` 选项强制删除  
`git rm --cached readme.txt` 删除readme.txt在仓库中的跟踪，**只是跟踪，但不删除文件**  
`git rm log/\*.log` 可以使用glob模式匹配，注意到星号 * 之前的反斜杠 \ 可以递归删除  
`git rm --cached -r *depend/* ` 递归删除所有depend目录下的所有问题  不论depend是否在根目录
 

### 1.9 移动文件 git mv

`git mv README.txt README` 相当于 `mv README.txt README`  `git rm README.txt`  `git add README` 三条命令

### 1.10 查看提交历史 git log

`git log` 会按照提交时间把所有提交记录都列出来   
选项： 
* `-p`显示提交两次之间的差异   
* `-2`最近两次的提交记录
* `--stat` 简单的增改行统计  
* `--pretty` 指定显示形式 `git log --pretty=oneline`:显示在一行 `git log --pretty=format:"%h - %an, %ar : %s"`:指定显示格式  
    选项 说明  
    %H 提交对象（commit）的完整哈希字串  
    %h 提交对象的简短哈希字串   
    %T 树对象（tree）的完整哈希字串  
    %t 树对象的简短哈希字串   
    %P 父对象（parent）的完整哈希字串   
    %p 父对象的简短哈希字串    
    %an 作者（author）的名字
    %ae 作者的电子邮件地址   
    %ad 作者修订日期（可以用 -date= 选项定制格式）   
    %ar 作者修订日期，按多久以前的方式显示   
    %cn 提交者(committer)的名字  
    %ce 提交者的电子邮件地址   
    %cd 提交日期   
    %cr 提交日期，按多久以前的方式显示   
    %s 提交说明   
可是使用图像化窗口查看提交记录会更方便，gitk则是一个图形化的工具，可以查看所有选项

### 1.11 撤销操作 

`git commit --amend` 撤销最后一次提交记录   
如：上次忘记添加forgotten_file,撤销后再提交

```shell
git commit -m 'initial commit'
git add forgotten_file
git commit --amend
```
:question:撤销最后一次是从库里又放回了暂存区？或者放回了已修改。抑或把之前的都弄不见了

`git reset HEAD benchmarks.rb` 取消暂存，如两个文件需要提交两次，但已经全都暂存了    
`git checkout -- benchmarks.rb` 取消修改，改变后又想回到原来的，:heart: 切记之前的修改将再也找不到了

### 1.12远程仓库操作 git remote
:question:远程仓库是什么概念？

`git remote` 查看当前配置有哪些远程仓库,列出每个远程库的简短名字,至少可以看到一个名为 origin 的远程库默认使用这个名字来标识你所克隆的原始仓库   
`git remote -v` -v = —verbose 显示对应的克隆地址   
`git remote  add [shortname] [url]` 添加远程仓库   
`git fetch [remote-name]` 从远程仓库抓取数据，此命令会到远程仓库中拉取所有你本地仓库中还没有的数据   
`git fetch origin`  会抓取从你上次克隆以来别人上传到此远程仓库中的所有更新    
`git push [remote-name] [branch-name]` 如果要把本地的 master 分支推送到 origin 服务器上（再次说明下，克隆操作会自动使用默认的master 和 origin 名字），可以运行下面的命令 `git push origin master`   
`git remote show origin` git remote show [remote-name] 查看某个远程仓库的详细信息   
`git remote rename pb paul`修改某个远程仓库的简短名称把pb改为paul，对远程仓库的重命名，也会使对应的分支名称发生变化，原来的 pb/master 分支现在成了 paul/master   
`git remote rm paul` 移除对应的远端仓库paul   

### 1.13 打标签 git tags
>对某一时间点上的版本打上标签,如对要发布的某个版本打上标签

`git tag` 列出现有标签,标签按字母顺序排列   
`git tag -l 'v1.4.2.*'` v1.4.2.1、v1.4.2.2、v1.4.2.3 …… 搜索相关标签   
`git tag -a v1.4 -m 'my version 1.4'` 新建含附注的标签   
`git tag -s v1.5 -m 'my signed 1.5 tag` 签署标签   
`git tag v1.4-lw` 新建轻量级标签    
`git push origin v1.5` 分享标签，默认push不会标签传输到远端服务器  
`git push origin --tags` 一次推送所有（本地新增的）标签上去，可以使用 `--tags` 选项  

### 1.14 小技巧

* 自动完成
    使用bash脚本`git-completion.bash`
* Git 命令别名
```shell
    git config --global alias.co checkout
    git config --global alias.br branch
    git config --global alias.ci commit
    git config --global alias.st status
```

## 二、Git分支 
>Git 中的分支，其实本质上仅仅是个指向 commit 对象的可变指针。Git会使用 master 作为分支的默认名字。在若干次提交后，你其实已经有了一个指向最后一次提交对象的 master 分支，它在每次提交的时候都会自动向前移动。

:question: 文件快照怎么会事，为何不用冗余文件，更轻量。分支指针，又是什么概念？

### 2.1 创建分支  branch
`git branch testing` 创建一个testing分支，但并不会切换到新的分支，还在master分支上   
`git checkout testing` 切换到testing分支,每次提交时都记录了祖先信息

![分支创建](image/git_branch.png)

master 分支指针指向一个版本地址   
testing 分支每提交一次分支指针都会向前移动    
head 指针指向当前正在工作的分支和版本  

### 2.2 基本的分支与合并  

假设你正在改问题issue53，突然接到一个电话说有个很严重的问题需要紧急修补，那么可以按照下面的方式处理：  
1. 返回到原先已经发布到生产服务器上的分支。  
2. 为这次紧急修补建立一个新分支。   
3. 测试通过后，将此修补分支合并，再推送到生产服务器上。   
4. 切换到之前实现新需求的分支，继续工作。   

git处理此问题的过程

1. 建立分支  `git checkout -b iss53` -b参数相当于 `git branch iss53` 和  `git checkout iss53` 
2. 修改#iss53问题  
3. 切换master分支 
     转换分支的时候最好保持一个清洁的工作区域。稍后会介绍几个绕过这种问题的办法（分别叫做 stashing 和 amending）。目前已经提交了所有的修改，所以接下来可以正常转换到 master 分支： `git checkout master` 
4. 创建紧急修复分支 `git checkout -b 'hotfix` 要测试确保成功
5. 合并hotfix分支到master分支 `git checkout master` `git merge hotfix` 
    合并是提示`Fast forward”（快进）`说明主分支指针只是向前移动而已，这是因为他们之间没有分歧，是顺序延伸下来的。  
    ![合并分支前](image/git_merge1.png)  ![合并分支后](image/git_merge2.png)
6. 合并后之前的 hotfix 分支将不再有用，可以删除 `git branch -d hotfix`
7. 回到之前#53问题分支继续修改问题 `git checkout iss53` 之后可以将master分支合并到iss53分支 `git merge master` 也可以等iss53修改完成后回到master`git checkout master` 将iss53分支合并到master分支 `git merge iss53`   
![合并分支前](image/git_merge3.png)  ![合并分支后](image/git_merge4.png)    
![新生成的commit](image/git_merge5.png) 
8. iss53没有了可以被删除 `git branch -d iss53`

### 2.3 冲突解决 
>有时候合并操作并不会如此顺利。如果你修改了两个待合并分支里同一个文件的同一部分，Git 就无法干净地把两者合到一起（译注：逻辑上说，这种问题只能由人来解决）。如果你在解决问题 #53 的过程中修改了 hotfix 中修改的部分，将得到类似下面的结果：

```shell
git merge iss53
Auto-merging index.html
CONFLICT (content): Merge conflict in index.html
Automatic merge failed; fix conflicts and then commit the result.
```
1. 查看冲突`git status` 查看冲突 任何包含未解决冲突的文件都会以未合并（unmerged）状态列出
2. 手动解决冲突  :question:将其中一个内容合并到另一个文件中，最初的文件删掉？
3. 标记为已解决(resolved) `git add` 一旦暂存，就表示冲突已经解决,如果你想用一个有图形界面的工具来解决这些问题，不妨运行 git mergetool，它会调用一个可视化的合并工具并引导你解决所有冲突
4. 如果觉得满意了，并且确认所有冲突都已解决，也就是进入了暂存区，就可以用 `git commit` 来完成这次合并提交。提交的记录差不多是这样：
```shell
Merge branch 'iss53'
Conflicts:
index.html
# It looks like you may be committing a MERGE.
# If this is not correct, please remove the file
# .git/MERGE_HEAD
# and try again.
```

### 2.4 分支管理

`git branch` 命令不仅仅能创建和删除分支，如果不加任何参数，它会给出当前所有分支的清单

```shell
git branch
iss53
* master     # ‘*’表示当前所在分支
testing
```

`git branch -v：` 查看各分支最后一次commit信息

```shell
git branch -v
iss53 93b412c fix javascript issue
* master 7a98805 Merge branch 'iss53'
testing 782fd34 add scott to the author list in the readmes
```

可以用 `--merge` 和 `--no-merged` 查看已合并或未合并的分支

```shell
git branch --merged       #查看已经合并分支
iss53
* master
#一般来说，列表中没有 * 的分支通常都可以用 git branch -d 来删掉，已经合并了就没有存在的意义了

git branch --no-merged    #查看未合并分支
testing
#git branch -d testing 删除未合并的分支会导致失败  如果要强行删除可以用 -D 选项删除
```

### 2.5 分支式工作流程
> 由于分支特别的方便所有才有了这种分支式工作流程

**1、长期分支**  
工作中常保留三个分支：稳定分支(master)、开发分支(develop)、新特性分支(topic).为的是不同层次的稳定性

![流水式](image/git_topic.png)

**2、特性分支**   
一些短期的工作内容，功能单一，完成测试稳定后合入稳定分支  
![特性分支的提交历史](image/git_topic1.png)

### 2.6 远程分支 （remote branch）**   
>远程分支是对远程仓库状态的一个索引，我们用 (远程仓库名)/(分支名) 这样的形式表示远程分支。当本地分支修改后，远程分支的指针不会改变。其他人推送版本到远程分支后，可以用 `git fetch origin ` 获取远程分支最新数据到本地

:question: 远程分支和本地分支有冲突怎么处理

**1、推送**   
    `git push (远程仓库名) (分支名)` 推送到远程分支可与他人共享

**2、跟踪分支**  
    从远程分支检出的本地分支，称为跟踪分支(tracking branch) 

**3、删除远程分支**    
`git push [远程名]:[分支名]`删除远程分支
```shell
git push origin :serverfix
To git@github.com:schacon/simplegit.git
- [deleted] serverfix
```
### 2.7 衍合 （rebase） 
>把一个分支整合到另一个分支的办法有两种： merge（合并） 和 rebase（衍合）   

**1、衍合的基础**  
 
把在 C3 里产生的变化补丁重新在 C4 的基础上打一遍。在 Git 里，这种操作叫做衍合（rebase）   
```shell
git checkout experiment
git rebase master
First, rewinding head to replay your work on top of it...
Applying: added staged command
```
![分支衍合](image/git_rebase1.png)   ![master快进](image/git_rebase2.png)   
衍合在提交历史上看到是顺序发生的，闭merge跟清晰

**2、更多有趣的衍合**

可以在衍合分支以外的地方衍合   
![衍合演示](image/git_rebase3.png)

**3、衍合的风险**

`永远不要衍合那些已经推送到公共仓库的更新`    

:turtle: 现在还不能很好理解，也暂时用不到，日后再补上

## 三、服务器上Git

> **Git 服务器**  一个远程的稳定的共享Git仓库，远程仓库通常只是一个 纯仓库(bare repository) ——一个没有当前工作目录的仓库，简单的说，纯仓库是你项目里 .git 目录的内容，别无他物。

### 3.1 协议

**1、本地协议**

**2、SSH协议**

**3、Git协议**

**4、HTTP/S协议**

### 3.2 在服务器上部署Git
:turtle: 用的时候再学

### 3.3 Git托管服务

**1、GitHub**   
**2、建立账户**   
**3、建立新仓库**   
**4、从Subversion中导入写项目**   
**5、开始合作**   
**6、项目页面**   
**7、派生(Forking)项目**  
**8、GitHub总结**   

## 四、分布式Git

## 五、Git工具

## 六、自定义Git

## 七、Git与其他系统

## 八、Git内部原理

### 8.1 底层命令(Plumbing)和高层命令(porcelain)

`.git` 目录包含了整个仓库的内容，要备份或者复制库。复制这个目录即可。

```shell
HEAD            #文件指向当前分支
branches/
config          # 文件包含了项目特有的配置选项
description     # 文件仅供GitWeb 程序使用
hooks/          #目录包含了的客户端或服务端钩子脚本
index           #文件保存了暂存区域信息
info/
objects/        #目录存储所有数据内容
refs/           #目录存储指向数据 (分支) 的提交对象的指针
.gitignore      #文件中管理的忽略模式
```

### 8.2 Git对象
>Git 是一套内容寻址文件系统,这种说法的意思是，从内部来看，Git 是简单的 key-value 数据存储.它允许插入任意类型的内容，并会返回一个键值，通过该键值可以在任何时候再取出该内容。
