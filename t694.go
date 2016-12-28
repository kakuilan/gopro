//使用Select实现Channel超时机制
package main
import (
		"fmt"
		"time"
		)

func main(){
	ch := make(chan int)
	timeout := make(chan bool, 1)

	go func(){
		time.Sleep(5 * time.Second)
		timeout <- true
	}()

	select {
		case <- ch:
		case <- timeout:
			fmt.Println("timeout...")
			break
	}
}
