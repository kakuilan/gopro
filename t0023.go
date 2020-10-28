//使用锁在多个协程间保证数据的一致性
package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		wg      sync.WaitGroup
		numbers []int
		mux     sync.Mutex
	)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			mux.Lock()
			numbers = append(numbers, i)
			mux.Unlock()
			wg.Done()
		}(i)

	}

	wg.Wait()
	fmt.Println("The number is:", numbers)
}
