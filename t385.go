// 单向channel. 可将channel隐式转换为单向队列,只收或只发
package main
import "fmt"

func main(){
	c := make(chan int, 3)

	var send chan<- int = c //send-only
	var recv <-chan int = c //receive-only

	send <-1
	<-recv

	fmt.Println(c, send, recv)

}
