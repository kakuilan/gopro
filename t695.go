//time.After的超时机制
package main
import (
		"fmt"
		"time"
		)

func main(){
	ch := make(chan int)
	lock := make(chan bool)
	go func(){
		for{
			select {
				case <- ch:
				case <- time.After(5 * time.Second) :
					fmt.Println("timeout...")
					lock <- true
					break
			}
		}
	}()

	<- lock
}
