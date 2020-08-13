### Go 编译
1. 在项目目录下执行 `go build`
2. 在其他路径下执行 `go build`, 后面需要加上项目的路径(从GOPATH/src后开始写起) 编译后的可执行文件就保存在当前目录下
3. 使用 `go build -o hello` 命令给编译后的文件重命名