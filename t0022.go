//使用共享变量(内存)在多个goroutine操作,将会出现数据不一致的问题
package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		wg      sync.WaitGroup
		numbers []int
	)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			numbers = append(numbers, i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	//打印的结果有可能不会是完整的0~9
	fmt.Println("The numbers is:", numbers)

}
