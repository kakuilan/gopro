//Channel迭代操作
package main
import "fmt"

func main(){
	ch := make(chan int)
	go func(){//匿名函数
		for i:=0;i<5;i++{
			ch <- i
		}
		close(ch)
	}()		

	for value := range ch {//rang迭代器
		fmt.Println(value)
	}

}
