# Go 并发教程

## 目录

## 1. GoRoutine 协程
Go Routine 在执行上来说是轻量级的线程.  
使用 **go f(s)** 在一个Go Routine 中调用这个函数。  
    
    go f("goroutine")

当然可以使用匿名函数启动一个 Go Routine. 

    go func(msg string){
	    fmt.Println(msg)
	}("going")

这两个 Go Routine 会异步地运行，所以需要等待它们都执行完毕。
其输出结果为:

    goroutine : 0
    going
    goroutine : 1
    goroutine : 2

这种交替输出的情况表示 Go 运行时是以异步的方式运行 Routine 的。

