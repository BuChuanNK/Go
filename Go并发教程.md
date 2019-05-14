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

## 2. Channel 通道
*通道* 是连接多个 Go Routine 的管道。 
通过 Go Routine 将值发送到通道, 在另一个 Go Routine 接收。

可以使用 make(chan val-type) 创建一个新的通道. 通道类型就是所需要传递值的类型。

    messages := make(chan string)

使用 channel <- 语法 发送一个新的值到通道中。
使用 <-channel 语言从通道中 接受一个值。

    go func() { message <- "ping" }()
    msg := <-messages

默认发送和接受操作是阻塞的, 直到发送方和接收方都准备完毕。  

**通道缓冲**

默认通道是 *无缓冲* 的, 这意味着只有在对应的接收 (<-chan) 通道准备号接收时, 才允许进行发送 (chan <-).   

*可缓存通道* 允许在没有对应接收方的情况下, 缓存限定数量的值.   

    message := make(chan string, 2)

**通道同步**

可以使用 *通道* 来同步 Go Routine 间的执行状态。
这里是一个使用阻塞的接受方法来等待一个 Go Routine 的运行结束。

    // 通道同步
    // 接受done通道中的bool值，告诉这个routine，其他routine结束了。
    func worker(done chan bool) {
        fmt.Print("working...")
        time.Sleep(time.Second)
        fmt.Println("done")

        // 发送true到done通道，告诉别的routine，这个routine结束了。
        done <- true
    }

在主程序中, Main Go Routine, 需要接收到 worker 来自 done 通道的确认bool值, 才能结束主协程.   

**通道方向**

当使用通道作为函数的参数时，可以指定通道是不是只用来发送或者接受值。
这个特性可以提升程序的类型安全性。

ping 函数定了一个只允许发送数据的通道。

    func ping(pings chan<- string, msg string){
        pings <- msg
    }
pong 函数允许通道 pings 来接收数据, 另一通道 pongs 来发送数据。

    func pong(pings <-chan string, pongs chan<- string){
        msg := <-pings
        pongs <- msg
    }

**通道选择器**

Go 的 *通道选择器* 让你可以同时等待多个通道操作。
Go 协程和通道以及选择器的结合是Go的一个强大特性。

使用 select 关键字来同时等待这两个值，并打印各自接收到的值。

    for i:=0; i<2; i++{
		select{
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}