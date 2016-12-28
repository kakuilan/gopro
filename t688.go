//Goroutine和Channel使用例子
package main
import (
		"fmt"
		"math/rand"
		)

func Test(ch chan int) {
	ch <- rand.Int() //向Channel中写入一个随机数
	fmt.Println("Go...")
}

func main(){
	chs := make([]chan int, 10)
	for i:=0;i<10;i++{
		chs[i] = make(chan int)
		fmt.Println(i)
		go Test(chs[i]) //启动10个Goroutine
	}

	for _,ch := range chs{
		value := <- ch //堵塞等待退出信号
		fmt.Println(value)
	}

}
