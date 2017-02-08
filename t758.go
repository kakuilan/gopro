//go timer
package main

import (
	"fmt"
	"time"
)

func main() {
	//初始化通道
	chl1 := make(chan int, 1000)
	sign := make(chan byte, 1)

	//给chl1通道写入数据
	for i := 0; i < 1000; i++ {
		chl1 <- i
	}

	//单独起一个groutine执行select
	go func() {
		var e int
		ok := true
		//首先声明一个*time.Timer类型的值,然后在相关case之后声明的匿名函数中尽可能的复用它
		var timer *time.Timer
		for {
			select {
			case e = <-chl1:
				fmt.Printf("chl1 -> %d\n", e)
			case <-func() <-chan time.Time {
				if timer == nil {
					//初始化到期时间据此间隔1ms的定时器
					timer = time.NewTimer( time.Millisecond)
				} else {
					//复用,通过Reset方法重置定时器
					timer.Reset( time.Millisecond)
				}
				//得知定时器到期事件来临时,返回结果
				return timer.C
			}():
				fmt.Println("timeout.")
				ok = false
				break
			}
			//终止for循环
			if !ok {
				sign <- 0
				break
			}
		}
	}()

	//惯用手法,读取sign通道数据,为了等待select的goroutine执行
	<-sign

}
