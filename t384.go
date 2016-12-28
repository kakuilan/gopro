// 内置函数len返回未被读取的缓冲元素数量,cap返回缓冲区大小
package main
import "fmt"

func main(){
	d1 := make(chan int)
	d2 := make(chan int, 3)

	d2 <- 1

	fmt.Println(len(d1), cap(d1)) // 0 0
	fmt.Println(len(d2), cap(d2)) // 1 3


}
