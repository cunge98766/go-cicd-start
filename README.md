# go-cicd-start

没有基础，直接干
前言：这是在window环境下开发的，后补上是因为有些命令执行的斜杠不一样，我的命令可能需要根据操作系统进行修改
----------思路----------------
cicd基本流程
1.从git中拉取代码，需要输入git地址，账号，密码，分支
2.将代码通过mvn打包成jar
3.通过dockerfile构建成docker镜像
4.将docker镜像打包上传到harbor镜像仓库  需要镜像仓库地址，账号，密码
5.写一套k8s delpoyment、service文件，然后将镜像地址替换成本次生成的地址
6.通过kubectl命令生成相关服务，达到发版部署的目的  你需要能连接到相关k8s地址，并有通过命令操作的权限

当然这只是初步的思路，后续做的时候，可能继续改进

第一步，学习go语言，
----------项目命令------------
自己做记录，要不经常忘记

项目启动
go run .\cmd\main.go
修改注释后生成新的api
.\swag.exe init -g .\cmd\main.go
刷新mod文件，解决各种不知名问题，时间上一个没解决，都是我自己名字打错了
go mod tidy

----------20240510----搭建一个go服务--------
准备使用go语言搭建一个ci cd功能的框架

虽然没有基础，但是我喜欢直接实践，通过实践的过程学习go

第一天创建了一个hello go文件
当拿到一个go程序，搜索package main就可以找到对应的初始方法

go mod init GO-CICD-START
这个命令是初始化项目，他的作用后续你用到的包都会在生成的go.mod中列出来，具体作用后面再说

--------20240511-----使用gin创建一个可以访问的go后台服务，并带swag openapi-------
首先需要创建一个可以一直运行的服务，就是用gin了，因为使用起来很简单

使用命令
go get -u github.com/gin-gonic/gin
就可以下载gin到本地
使用
 ginService := gin.Default()
 //各种get post方法
 ginService.Run(":8082")
创建一个服务，如此简单

//gin 可以和swag一起用，
这是地址：<https://github.com/swaggo/gin-swagger>
这里面有教程
第一步执行 go版本1.17以上哦，要不用get
go install github.com/swaggo/swag/cmd/swag@latest
注意你通过这个命令下载的东西，不一定在你的工程里面，用 go env
这个命令查看  set GOPATH=D:\GOwork
set GOPATH=后面就是你的下载路径，这个路径的bin文件夹下面就会出现 swag.exe,复制到你的工程就可以了
./swag init  理论执行就可以了
实际上有些问题，由于我的main方法不在根目录下，在cmd下，经测试，需要执行
.\swag.exe init -g .\cmd\main.go
当你有各种报错，试试这个命令可能会有用
go mod tidy
这步之后，我已经可以启动我得go服务，并生成对应的swag，并能通过swag访问我本地的接口获取

这玩意真次，不想用了，完全依赖注释，注释错了，就有可能不能访问，都不能自动识别，我要你干嘛
但还是记录一下

swag的  注释解释 不解释了，测了几个发现太死板了，写个标题和路径得了
// @BasePath /api/v1
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example                    ------我是标题
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /json [get]           ------我是路径 get/post方法都不能自动识别，注释写错了，生成就错了

我就用了这三个，也满足我现有的，太辣鸡，不能自动识别，注释写错了，就生成错了，也不能根据代码自动生成docs，
// @Tags 标题
// @Router /json [get]    路径   []里面是get/post方法类型
// @Description 描述，我是干嘛的

-----------20240513-------------------------------
打log，真烦，日志中需要增加log，我将采用go标准库自带的 log/slog
使用这个slog 需要go版本在1.21以上
今天有事，先不学了
--------------------------------------------
