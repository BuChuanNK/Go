# Go 语言基础
## Go 简介
> Go 由 Google 开发，用于建立简单，快速，可靠软件的程序设计语言
> 
> 参考书籍：《Go by Example》

## 1. 你好，世界
HelloWorld测试代码：

    package main
    
    import "fmt"
    func main(){
        fmt.Println("hello world!")
    }

如何运行HelloWorld.go文件：

    直接运行
    $ go run HelloWorld.go
    编译为二进制文件
    $ go build HelloWorld.go
    $ ./HelloWorld

## 2. 值
Go 拥有的值类型，包括String ,int float, boolean

Var.go测试代码：

    package main

    import "fmt"
    func main(){
        fmt.Println("go"+"lang")

        fmt.Println("1+1=", 1+1)
        fmt.Println("7.0/3.0 = ", 7.0/3.0)

        fmt.Println(true && false)
        fmt.Println(true || false)
        fmt.Println(!true)
    }


## 3. 变量
使用var来显式声明一个变量。