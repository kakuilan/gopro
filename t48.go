package main
import "fmt"

func main(){
    var ch1 chan int //ch1是一个正常的channel,不是单向的
    var ch2 chan<- float64 //ch2是单向的channel,只用于写float64数据
    var ch3 <-chan int	//ch3是单向channel,只用于读取int数据
    ch4 := make(chan int)
    ch5 := <-chan int(ch4) //ch5就是一个单向的读取channel
    ch6 := chan<- int(ch4) //ch6是一个单向的写入channel
    fmt.Println(ch1, ch2, ch3, ch4, ch5, ch6)
}
