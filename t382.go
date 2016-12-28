// 异步channel
// 通常情况下,异步channel可减少排队阻塞,具备更高的效率.但应该考虑使用指针规避大对象拷贝,将多个元素打包,减少缓冲区大小等.
package main
import "fmt"

func main(){
	data := make(chan int, 3) //缓冲区可以存储3个元素
	exit := make(chan bool)

	data <- 1 //在缓冲区未满前,不会堵塞
	data <- 2
	data <- 3

	go func(){
		for d := range data { //在缓冲区未空前,不会堵塞
			fmt.Println(d)		
		}
		
		exit <- true
	}()

	data <- 4 //如果缓冲区已满,堵塞
	data <- 5
	close(data)

	<-exit

}

