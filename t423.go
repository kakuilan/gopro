// Data Race 数据竞争
// go run -race数据竞争检测会严重影响性能
package main
import "sync"
import "fmt"

func main(){
	var wg sync.WaitGroup
	wg.Add(2)

	x:= 100

	go func(){
		defer wg.Done()

		for {
			x +=1
		}
	}()

	go func(){
		defer wg.Done()
		for {
			x +=100
		}
	}()

	wg.Wait()
	fmt.Println(x)
}
