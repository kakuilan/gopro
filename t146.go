//多线程
package main
import (
    "fmt"
    "runtime"
    "time"
)

func SayHello(){
    for i:=0;i<10;i++{
	fmt.Println("Hello ")
	//让出时间片
	runtime.Gosched()
    }
}

func SayWorld(){
    for i:=0;i<10;i++{
	fmt.Println("World!")
	runtime.Gosched()
    }    
}

func main(){
    go SayHello()
    go SayWorld()
    time.Sleep(5 * time.Second)
}
