//单向Channel
package main
import "fmt"

func Recv(ch <- chan int, lock chan<- bool) {
	for value:= range ch{
		fmt.Println(value)
	}
	lock <- true
}

func Send(ch chan<-int) {
	for i:=0;i<5;i++{
		ch <- i
	}
	close(ch)
}

func main(){
	ch := make(chan int) //双向Channel,可转换为单向Channel
	lock := make(chan bool)
	go Recv(ch, lock) //只能从ch接收的Goroutine
	go Send(ch) //只能向ch发送的Goroutine
	<-lock
}
