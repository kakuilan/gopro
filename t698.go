//强行终止goroutine
package main
import (
		"fmt"
		"time"
		"runtime"
		)

func main(){
	go func(){
		defer fmt.Println("Goroutine1 defer...")
		for i:=0;i<10;i++{
			if i==5 {
				runtime.Goexit()
			}
			fmt.Println("Goroutine1: ", i)
		}	
	}()

	go func(){
		fmt.Println("Goroutine2")
	}()

	time.Sleep(5 * time.Second)
}
