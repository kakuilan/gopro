// channel选择.
// 如果需要同时处理多个channel,可使用select语句.它随机选择一个可用channel做收发操作,或执行default case.
package main
import "fmt"
import "os"

func main(){
	a,b := make(chan int, 3), make(chan int)
	
	go func(){
		v,ok,s := 0,false, ""
		for {
			select { //随机选择可用channel,接收数据
				case v,ok = <-a :
					s = "a"
				case v,ok = <-b :
					s = "b"
			}

			if ok {
				fmt.Println(s,v)
			}else{
				os.Exit(0)
			}
		}	
	}()

	for i:=0;i<5;i++{ //随机选择可用channel,发送数据
		select {
			case a <- i:
			case b <- i:
		}
	}

	close(a)
	select{} //没有可用channel,堵塞main goroutine



}
