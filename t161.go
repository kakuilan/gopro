// time.AfterFunc
package main
import (
    "fmt"
    "time"
)

func Test(){
    fmt.Println("Hello world!", time.Now())
}

func main(){
    fmt.Println(time.Now())
    time.AfterFunc(10 * time.Second, Test)
    var str string
    fmt.Scan(&str) //等待用户输入,不让进程结束,否则进程结束定时器也就无效了

}
