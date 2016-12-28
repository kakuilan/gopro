// channel 单向通信
package main
import "fmt"

//chan<- int仅能发送
func producer(c chan<- int){
    defer close(c) //关闭channel
    for i:=0;i<10;i++{
	c <- i //堵塞,直到数据被消费者取走后才能发送下一条数据
    }
}

//<-chan int仅能接收
func consumer(c <-chan int, f chan<- int) {
    for v:= range c {
	fmt.Println(v)
    }
    f <- 1 //向F发送一个数据,告诉main数据已接收完成
}

func main(){
    buf := make(chan int)
    flg := make(chan int)
    go producer(buf)
    go consumer(buf,flg)
    <-flg //等待数据接收完成
}
