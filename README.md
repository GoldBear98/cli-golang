# cli-golang
**1.概述**
命令行实用程序并不是都象 cat、more、grep 是简单命令。go 项目管理程序，类似 java 项目管理 maven、Nodejs项目管理程序npm、git 命令行客户端、 docker 与 kubernetes 容器管理工具等等都是采用了较复杂的命令行。即一个实用程序同时支持多个子命令，每个子命令有各自独立的参数，命令之间可能存在共享的代码或逻辑，同时随着产品的发展，这些命令可能发生功能化、添加新命令等。因此，符合 OCP原则 的设计是至关重要的编程需求。

任务目标：
1.熟悉 go 命令行工具管理项目
2.综合使用 go 的函数、数据结构与接口，编写一个简单命令行应用 agenda
3.使用面向对象的思想设计程序，使得程序具有良好的结构命令，并能方便修改、扩展新的命令,不会影响其他命令的代码
4.项目部署在 Github 上，合适多人协作，特别是代码归并
5.支持日志（原则上不使用debug调试程序）

**2、GO 命令**
2.1 go 命令的格式
使用：

```
go command [arguments]
```
版本（go 1.8）的命令有：

```
build       compile packages and dependencies
clean       remove object files
doc         show documentation for package or symbol
env         print Go environment information
bug         start a bug report
fix         run go tool fix on packages
fmt         run gofmt on package sources
generate    generate Go files by processing source
get         download and install packages and dependencies
install     compile and install packages and dependencies
list        list packages
run         compile and run Go program
test        test packages
tool        run specified go tool
version     print Go version
vet         run go tool vet on packages
```
2.2 go 命令分类
1.环境显示： version、env
2.构建流水线： clean、build、test、run、（publish/git）、get、install
3.包管理： list, get, install
4.杂项：fmt，vet，doc，tools …
具体命令格式与参数使用 ：

```
go help [topic]
```
**3.准备知识或资源**
3.1 Golang 知识整理
这里推荐 time-track 的个人博客，它的学习轨迹与课程要求基本一致。以下是他语言学习的笔记，可用于语言快速浏览与参考：
[《Go程序设计语言》要点总结——程序结构](https://niyanchun.com/)
[《Go程序设计语言》要点总结——数据类型](https://niyanchun.com/)
[《Go程序设计语言》要点总结——函数](https://niyanchun.com/)
[《Go程序设计语言》要点总结——方法](https://niyanchun.com/)
[《Go程序设计语言》要点总结——接口](https://niyanchun.com/)
以上仅代表作者观点，部分内容是不准确的，请用批判的态度看待网上博客。切记：
·GO 不是面向对象(~~OOP~~ ) 的。 所谓方法只是一种[语法糖](https://baike.baidu.com/item/%E8%AF%AD%E6%B3%95%E7%B3%96)，它是特定类型上定义的操作（operation）
·指针是没有 nil 的，这可以避免一些尴尬。 p.X 与 v.x (p 指针， v 值) 在语义上是无区别的，但实现上是有区别的 p.x 是实现 c 语言 p->x 的语法糖
·zero 值好重要

3.2JSON 序列化与反序列化
参考：[JSON and Go](https://blog.go-zh.org/json-and-go)
json 包是内置支持的，文档位置：https://go-zh.org/pkg/encoding/json/

3.3 复杂命令行的处理
这里我们选择cobra这个工具
使用命令`go get -v github.com/spf13/cobra/cobra`下载过程中会提示错误，解决办法是：
在`$GOPATH/src/golang.org/x`目录下用`git clone`下载sys和text项目，具体指令为：

```
git clone https://github.com/golang/sys
```

```
git clone https://github.com/golang/text
```
然后使用`go install github.com/spf13/cobra/cobra` ，安装后在 $GOBIN 下出现了 cobra 可执行程序。

Cobra 的简单使用
创建一个处理命令`agenda register -uTestUser `或`agenda register --user=TestUser`的小程序。
简单步骤如下：

```
cobra init
cobra add register
```
需要的文件就产生了。 你需要阅读 main.go 的 main() ; root.go 的 Execute();最后修改 register.go, init() 添加：

```
registerCmd.Flags().StringP("user", "u", "Anonymous", "Help message for username")
```
Run 匿名回调函数中添加：

```
username, _ := cmd.Flags().GetString("user")
fmt.Println("register called by " + username)
```
测试命令：

```
$ go run main.go register --user=TestUser
register called by TestUser
```
**4.agenda 开发项目**
Agenda业务需求
用户注册：
注册新用户时，用户需设置一个唯一的用户名和一个密码。另外，还需登记邮箱及电话信息。如果注册时提供的用户名已由其他用户使用，应反馈一个适当的出错信息；成功注册后，亦应反馈一个成功注册的信息。

用户登录：
用户使用用户名和密码登录 Agenda 系统。用户名和密码同时正确则登录成功并反馈一个成功登录的信息。否则，登录失败并反馈一个失败登录的信息。

用户登出：
已登录的用户登出系统后，只能使用用户注册和用户登录功能。

agenda界面

```
./agenda
```
界面如下：

```
A meeting management system base on CLI
Usage:
     agenda [command]
 
 Available Commands:
     log_in                     log in
     log_out                   log out
     register                   Register a new user

Flags:
	  --config string   config file (default is $HOME/.agenda.yaml)	  
     -h, --help            help for agenda
     -t, --toggle          Help message for toggle

Use "agenda [command] --help" for more information about a command.
```

测试register：
创建用户Tony

```
agenda register -n [username] -p [password] -e [email] -t [phone]
```

```
agenda register -n Tony -p 123456 -e Tony@163.com -t 13246820667
Regist successfully!
```

登录账户：

```
agenda log_in -n [username] -p [password]
```

```
agenda log_in -n Tony -p 123456
Log in successfully!
```
退出账户就不在这里进行展示，至此也就完成了一个agenda的注册、登录、退出功能。
