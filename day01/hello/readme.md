### Go 命令

#### go build
- 在项目目录下执行 `go build`
-  在其他路径下执行 `go build go-study/day01/hello`, 后面需要加上项目的路径(从GOPATH/src后开始写起) 编译后的可执行文件就保存在当前目录下
-  使用 `go build -o hello` 命令给编译后的文件重命名
-  在项目上线的时候直接在本地打包放在服务器上就可以了, 服务器上不需要部署Go环境

#### go run 
-  `go run` 像执行脚本文件一样执行Go代码
#### go install 
-  `go install` 命令做了两步事情, 1.先编译得到一个可执行文件 2.将可执行文件拷贝到 GOPATH/bin

### Go 跨平台编译
例如: 在mac平台编译一个能在windows平台执行的可执行文件
```
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```
在mac平台编译一个能在linux平台执行的可执行文件
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
```
在windows编一个能在mac平台执行的可执行文件
```
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build
```
