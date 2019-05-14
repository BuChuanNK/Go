package main

import "fmt"
import "time"

// 通道同步
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

// 通道方向
func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {

	// 通道
	messages := make(chan string)

	go func() { message <- "ping" }()

	msg := <-messages
	fmt.Println(msg)

	// 通道缓冲
	message := make(chan string, 2)

	message <- "buffered"
	message <- "channel"

	fmt.Println(<-message)
	fmt.Println(<-message)

	// 通道同步
	done := make(chan bool, 1)
	go worker(done)

	<-done

	// 通道方向
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

	// 通道选择器
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
