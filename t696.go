//让出时间片
package main
import (
		"fmt"
		"time"
		"runtime"
		)

func main(){
	go func(){
		for i:=0;i<5;i++{
			if i==2 {
				runtime.Gosched() //出让时间片给下面的goroutine2
			}
			fmt.Println("Goroutine1:", i)
		}
	}()

	go func(){
		fmt.Println("Goroutine2")	
	}()

	time.Sleep(5 * time.Second)

}


