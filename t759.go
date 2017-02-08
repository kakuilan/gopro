//time.After函数,表示多少时间之后,但是在取出channel内容之前不阻塞,后续程序可以继续执行
package main
import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan int, 1)
    ch2 := make(chan int, 1)

    select {
    case e1 := <-ch1:
        //如果ch1通道成功读取数据,则执行该case
        fmt.Printf("1th case is selected. e1=%v", e1)
    case e2:= <-ch2:
        fmt.Printf("2th case is selected. e2=%v", e2)
    case <- time.After(2 *time.Second) :
        fmt.Println("Time out")
    }


}
