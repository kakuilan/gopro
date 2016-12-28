//除range之外,接收方还可使用for方式接收数据,并判断Channel是否关闭
package main
import "fmt"

func main(){
	ch := make(chan int)
	go func(){
		for i:=0;i<5;i++{
			ch <- i
		}
		close(ch)
	}()

	for {
		if value,ok := <-ch; ok{
			fmt.Println(value)
		}else{
			fmt.Println("channel is closed.")
			break
		}
	}

}
