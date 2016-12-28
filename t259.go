// 添加额外的退出channel
package main
import "fmt"

func main(){
    ch := make(chan int)
    quit := make(chan bool)
    go shower(ch, quit)
    for i:=0;i<10;i++{
	ch<- i
    }
    quit<- false
}

func shower(c chan int, quit chan bool) {
    for {
	select {
	case j := <-c :
	    fmt.Printf("%d\n", j)
	case <-quit: //从退出channel读取并丢弃该值
	    break
	}
    }

}
