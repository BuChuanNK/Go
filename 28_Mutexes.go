package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	// state 是一个 map
	var state = make(map[int]int)

	// 这里互斥锁 mutex 将同步对 state 的访问
	var mutex = &sync.Mutex{}

	// ops 将记录我们对 state 的操作次数
	var ops int64 = 0

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

	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
