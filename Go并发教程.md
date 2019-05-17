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

## 3. Timeout 超时处理
*超时* 对于一个连接外部资源, 或者其它一些需要花费执行时间的操作的程序而言很重要. 在Go语言当中，使用 Chan 和 select 就可以实现超时操作. 

    c1 := make(chan string, 1)
    // 执行一个匿名函数，在2秒内通过通道c1返回执行结果。
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()
    // 使用select实现一个超时操作。通道选择器中有两个选择:
    // 1. res := <-c1 等待结果
    // 2. <-Time.After 等待超时时间1秒后，自动执行结果。
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

## 4. Non-Blocking Channel 非阻塞通道操作
常规的通道发送和接收数据是阻塞的. 
然而, 外面可以使用带一个default子句的select来实现 *非阻塞* 的发送、接收,甚至是非阻塞的多路select。

    // 这是一个典型的非阻塞接收的例子. 如果在messages通道中存在就接收，不存在就执行default.
    select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

    // 使用通道选择器来实现多路的非阻塞通道操作。关键是default。
    select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

## 5. Closing Channel 通道关闭
*通道关闭* 意味着不能再向这个通道发送数据。
这可以用来给这个通道的接收方传达工作已经完成的信息。

    
    // 接收方
    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    // 发送方
    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    // 发送结束，关闭通道jobs
    close(jobs)
    // 此时jobs通道已经关闭，等到done通道的bool信息传送出来，表示接收方的工作已经完成。
    // 通道同步知识, 参见18_Channel.go
    <-done

## 6. Range Over Channels 通道遍历
可以使用for 和range的语法来遍历从通道中取得的所有值。
即遍历通道缓存中的所有值。

    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)

    for elem := range queue {
        fmt.Println(elem)
    }

## 7. Timer 定时器
Go 内置有 *Timer* 和 *Ticker*。
这两个特性可以实现定时任务和时间段内的循环任务。

定时器表示在未来某一时刻的独立事件。因此: 
1. 需要告诉定时器需要等待的时间。
2. 提供一个用于通知的通道。

    timer1 := time.NewTimer(time.Second * 2)

<-timer1.C 意味着: 直到这个定时器的通道C明确地发送了定时器失效的值之前，将一直阻塞. 

定时器可以在中途取消。

    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 expired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }

## 8. Ticker 打点器
*定时器* 是希望在未来某一刻执行一次时使用的。
*打点器* 时希望在固定的时间间隔重复执行时使用的。

打点器和定时器的机制有点相似: 
1. 需要一个通道来发送数据。
2. 使用内置的 range 来迭代每间隔时间发送数据。
   
    ticker := time.NewTicker(time.Millisecond * 500)
    go func() {
        for t := range ticker.C {
            fmt.Println("Tick at", t)
        }
    }()

打点器同样可以停止。但是停止之后就不能再接收到值来。
    
    time.Sleep(time.Millisecond * 1600)
    ticker.Stop()

## 9. Worker Pool 工作池
在 Go 中, 可以是使用 Go Routine 和 Channel 来实现一个工作池. 
*工作池* 可以用来实现高并发的任务。比如例子 25_WorkerPool.go 中, 每个worker都会从jobs通道接收任务，然后通过results发送出去。每个任务间隔为1s。

    // 声明3个worker
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }
    // 输入9个job
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)
    // 等待结果输出
    for a := 1; a <= 9; a++ {
        <-results
    }

最后结果会发现，9个任务只耗时3秒，因为3个worker是并行的。

## 10. Rate Limit 速率限制
*Rate Limit* 是一个重要的控制服务资源利用和质量的途径。
最简单的速率限制，就是先声明一个限制器，然后在for循环中必须要等到限制器时间到了，之后再进行操作，类似于Sleep()

    limiter := time.Tick(time.Millisecond * 200)
    for req := range requests {
        <-limiter
        fmt.Println("request", req, time.Now())
    }

但是，如果只是需要临时进行速率限制，并且不影响整体的速率，可以通过 *通道缓冲* 来实现。

    // 使用 通道缓冲 来实现3次临时的脉冲型速率限制
    burstyLimiter := make(chan time.Time, 3)
    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }
    // 每200ms将添加一个新值到burstyLimiter通道中，直到达到了3的限制。
    go func() {
        for t := range time.Tick(time.Millisecond * 200) {
            burstyLimiter <- t
        }
    }()
    // 当接入请求超过3个时，刚开始的3个将要受到Limiter的脉冲影响。
    burstyRequests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)
    for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }

## 11. Atom Counter 原子计数器
Go 中最主要的状态管理方法时通过 *通道间的沟通来完成的*。
还有一种方法时使用 *Sync/atomic* 包中, 在多个 Go Routine 中进行 *原子计数*。

可以使用无符号整数型数来表示计数器 (Uint64). 使用 AddUint64 来让计数器自动增加, 使用 & 语法来给出 ops 的内存地址. 

    // 自动计数, 同时使用 &ops 来输出内存地址
    atomic.AddUint64(&ops, 1)
    // 允许其他 Go Routine 执行
    runtime.Gosched()

通过 LoadUint64 将当前值的拷贝提取到 opsFinal 中。

    opsFinal := atomic.LoadUint64(&ops)

## 12. Mutexes 互斥锁
在 Go 中，对于更为复杂的多 Go Routine 场景，可以使用 *互斥锁* 来安全地访问数据。

    for r := 0; r < 100; r++ {
		go func() {
			total := 0
			// 每次循环读取, 使用一个键来进行访问,
			// Lock() 保证了这个mutex来确保对state的独占访问, 读取选定的键的值.
			// Unlock() 这个mutex, 并且 ops 值加1.
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)

				// 为了确保Go Routine不会在调度中饿死, 因此需要在每次操作后使用 runtime.Gosched()来进行释放.
				// 这个释放一般是自动处理, 例如每个通道操作后或者time.Sleep的阻塞调用后相似.
				runtime.Gosched()
			}
		}()
	}

## 13. Stateful GoRoutine Go状态协程
在Go中, 我们可以使用互斥锁来进行明确的锁定，使得让共享的 state 跨多个 Go Routine 同步访问. 
另一个选择是使用内置的 Go Routine 和 Channel 的同步特性来达到同样的效果. 