// 除了range外,还可用ok-idiom模式判断channel是否关闭
package main
import "fmt"

func main(){
	data := make(chan int, 3) //缓冲区可以存储3个元素
	exit := make(chan bool)

	data <- 1 //在缓冲区未满前,不会堵塞
	data <- 2
	data <- 3

	go func(){
		for {
			if d,ok := <-data;ok{
				fmt.Println(d)
			}else{
				fmt.Println("none")
				break
			}
		}	
		exit <- true
	}()

	data <- 4 //如果缓冲区已满,堵塞
	data <- 5
	close(data)

	<-exit

}

