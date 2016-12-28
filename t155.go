//select 超时
package main
import (
    "fmt"
    "time"
)

func main(){
    c := make(chan int)
    select {
	case <-c : //因为没有向c发送数据,所以会一直堵塞
	    fmt.Println("收到数据")
	case <-time.After(3 * time.Second) :
	    fmt.Println("超时退出")
    }

}
