// 用简单工厂模式打包并发任务和channel
package main
import "fmt"
import "os"

func NewConsumer() chan int {
	data := make(chan int, 3)
	
	go func(){
		for d := range data {
			fmt.Println(d)
		}

		os.Exit(0)
	}()

	return data
}

func main(){
	data := NewConsumer()

	data <- 1
	data <- 2

	close(data)
	select{}


}


