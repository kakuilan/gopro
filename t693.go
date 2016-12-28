//select机制
package main
import "fmt"

func main(){
	ch1 := make(chan int)
	ch2 := make(chan string)
	lock := make(chan bool)

	go func(){
		for {
			select {
			case value,ok := <-ch1:
				if !ok{
					lock <-true
					break
				}
				fmt.Println("ch1 = ", value)
			case value,ok := <-ch2:
				if !ok{
					lock <-true
					break
				}
				fmt.Println("ch1 = ", value)
			}//select	
			
		}
	}()

	ch1 <- 100
	ch2 <- "Golang"
	ch2 <- "go go go!"
	ch1 <- 250
	close(ch1)
	close(ch2)
	<-lock
}
