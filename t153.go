//channel 带缓冲
package main
import "fmt"

func main(){
    c := make(chan int, 2)
    c<- 1
    c<- 2 //此时若再向c发送数据,将会堵塞,运行时报错

    fmt.Println(<-c)
    fmt.Println(<-c) //此时若再从c中取数据,将出现堵塞,运行时报错
}
