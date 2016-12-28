// 和协程yield作用类似,Gosched让出底层线程,将当前goroutine暂停,放回队列等待下次被调度执行
package main
import (
		"sync"
		"runtime"
		)

func main(){
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func(){
		defer wg.Done()

		for i:=0;i<6;i++{
			println(i)
			if i==3 {
				runtime.Gosched()
			}
		}
	}()

	go func(){
		defer wg.Done()
		println("Hello, World!")
	}()

	wg.Wait()



}
