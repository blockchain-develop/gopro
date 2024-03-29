# git基础命令

## 安装

```
apt install git
```

## tag

创建tag:
git tag -a v1.0.0 -m "xxxxxxxx"
git tag -a v1.0.0 -m "version 1.0.0" 22860ed0e57465015b4a47c2f2d3be3d392e8dfc

推动tag:
git push origin v1.0.0

删除tag:
git tag -d v1.0.0

删除远程:
git push origin :refs/tags/v1.0.0

## 添加远程仓库&merge

git remote add origin(本地名字，默认为origin) [master(branch的名字，默认为master)] url

修改远程仓库地址

git remote set-url origin url

拉取remote最新代码
git fetch master  /  git pull origin master

push到remote

git push origin master

切换到另一个仓库的分支
git checkout -b poly poly/master

## 创建一个新分支并提交

创建本地分支
git checkout -b poly poly/master

提交到远端
git push poly poly

## 删除本地branch

git branch -D dev

## 删除远程branch

git push origin --delete dev

## add后撤销

git reset HEAD

撤销最后add的所有文件

## 如何添加本地项目到github

+ 在已存在的项目下，通过命令git init把这个文件夹变成Git可管理的仓库。里面多了个.git文件夹，它是Git用来跟踪和管理版本库的。
+ 通过git add把项目添加到仓库。在这个过程中你其实可以一直使用git status来查看你当前的状态。
+ 用git commit把项目提交到仓库。
+ 将Github上创建好Git仓库和本地仓库进行关联，在本地项目下执行命令：git remote add origin https://github.com/guyibang/TEST2.git
+ 把本地库的所有内容推送到远程仓库（也就是Github）上了，通过：git push origin master

## 从一个branch merge到另一个branch

假如我现在有两个branch：story6header, competition. 把competition这个branch merge到 story6header 这个branch 上。执行以下操作：
+ git checkout story6header #切换到story6header分支
+ git merge competition #将competition merge到story6header
+ git status #查看story6header上的更改
+ git add . #提交更改
+ git commit -m "add catagory function and seed" #提交commit信息
+ git push origin story6header

如果遇到问题：
1. 问题1

```
DELL@DESKTOP-KCAJBEQ MINGW64 ~/go/src/github.com/palettechain/explorer (main)
$ git merge master
fatal: refusing to merge unrelated histories
```

github默认不允许合并没有共同祖先的分支，需要加上 --allow-unrelated-histories。

```
DELL@DESKTOP-KCAJBEQ MINGW64 ~/go/src/github.com/palettechain/explorer (main)
$ git merge master --allow-unrelated-histories
Merge made by the 'recursive' strategy.
```

[git指导](https://www.cnblogs.com/yuqing-wei/p/5487713.html)

## 提交PR后出现冲突的解决办法

假设本地repository的branch A向另一个repository的branch master提交pr，出现冲突无法自动merge的解决办法。

+ 本地checkout并切换到branch A，pull拉取更新到最新代码
+ 在本地branch A上，merge远程repository的branch master
+ 会提示无法合并，手动解决完冲突并提交到branch A 
+ 回到PR，会发现PR已经无冲突
+ 让有merge权限的人进行merge即可

## 删除一个commit和push

+ 回到以前的一个点   git reset HEAD~1 --hard 
+ 再次修改并且commit
+ force的push，将以前的删除    git push -f wallet deriveAddress 

## 如何保存修改回退到以前版本之后再加进来

+ git stash
+ git stash pop
