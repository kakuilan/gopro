//分别使用锁和原子操作实现多个goroutine对同一变量进行累加操作
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var (
		mux    sync.Mutex
		mTotal = 0
		aTotal int64
	)

	for i := 0; i < 10; i++ {
		go func() {
			for {
				mux.Lock()
				mTotal += 1
				mux.Unlock()
				time.Sleep(time.Millisecond)
			}
		}()

		go func() {
			for {
				atomic.AddInt64(&aTotal, 1)
				time.Sleep(time.Millisecond)
			}
		}()

	}

	time.Sleep(time.Second)
	fmt.Println("The mTotal is:", mTotal)
	fmt.Println("The aTotal is:", aTotal)

}
