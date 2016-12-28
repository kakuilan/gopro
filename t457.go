// goroutine
package main
import (
		"fmt"
		"runtime"
		)


func say(s string) {
	for i:=0;i<5;i++{
		runtime.Gosched() //让出时间片
		fmt.Println(s)
	}
}

func main(){
	go say("world") //开一个新的goroutine
	say("hello") //当前goroutine执行

}



