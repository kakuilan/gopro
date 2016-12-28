// 可在初始化函数中使用goroutine,可等待其结束
package main
import (
		"fmt"
		"time"
		)

var now = time.Now()
	
func main(){
	fmt.Println("main:", int(time.Now().Sub(now).Seconds()))
}

func init(){
	fmt.Println("init:", int(time.Now().Sub(now).Seconds()))
	w := make(chan bool)

	go func(){
		time.Sleep(time.Second * 3)
		w <- true
	}()

	<-w

}
