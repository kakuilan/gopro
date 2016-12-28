// 调用runtime.Goexit将立即终止当前goroutine执行,调度器确保所有已注册defer延迟调用被执行
package main
import (
	"sync"
	"runtime"
)

func main(){
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func(){
		defer wg.Done()
		defer println("A.defer")
	
		func(){
			defer println("B.defer")
			runtime.Goexit() //终止当前goroutine
			println("B") //不会执行
		}()

		println("A") //不会执行
	}()

	wg.Wait()

}
