//channel超时
//有时会出现goroutine堵塞的情况,那么我们如何避免整个程序进入堵塞呢?\
//我们可以利用select来设置超时
package main
import "fmt"
import "time"

func main(){
	c := make(chan int)
	o := make(chan bool)

	go func(){
		for {
			select {
				case v := <-c :
						println(v)
				case <- time.After(5 * time.Second) :
						println("timeout")
						o <- true
						break
			}
		
		}
	}()

	t := <-o
	fmt.Println(t)

}
