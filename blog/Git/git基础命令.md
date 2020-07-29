# git基础命令

## tag

创建tag:
git tag -a v1.0.0 -m "xxxxxxxx"

推动tag:
git push origin v1.0.0

删除tag:
git tag -d v1.0.0

删除远程:
git push origin :refs/tags/v1.0.0

## 添加远程仓库&merge

git remote add origin(本地名字，默认为origin) master(branch的名字，默认为master) url

拉取remote最新代码

git pull origin master

push到remote

git push origin master

## add后撤销

git reset HEAD

撤销最后add的所有文件


[git指导](https://www.cnblogs.com/yuqing-wei/p/5487713.html)
