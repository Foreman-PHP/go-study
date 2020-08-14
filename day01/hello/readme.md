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
### Go 语言文件的基本结构
1. 所有的.go文件都需要有package关键字来声明当前文件属于哪个包
2. 如果当前 package 声明的是 main 包 说明最后代码会编译成可执行文件
3. 如果是工具类的包比如连接mysql的包是不需要编译成可执行文件的, 也就是不能声明为main
```
package main

// 导入语句 
import "fmt"

// 函数外只能放置标识符(变量/常量/函数/类型)的声明

// 程序的入库函数
func main() {
	fmt.Printf("hello world")
}
```

### 变量和常量

#### 标准声明
- 在函数体外声明为全局变量可以只声明不使用, 但是在函数内声明了必须要使用
```
var 变量名 变量类型
```

#### 批量声明
```
var(
 a string
 b int
 c bool
)
```

#### 简短变量声明
- 只能在函数内使用
```
name := "小明"__
```

#### 匿名变量
```
x, _ := foo()
```

#### 常量
```
// 单独声明
const STATUS = 200

// 批量声明
const (
    STATUSOk = 200
    NOTFOUND = 404
)
```

#### 常量计数器 iota