// time定时器 time.After
package main
import (
    "fmt"
    "time"
)

func main(){
    fmt.Println(time.Now())
    c := time.After(10 * time.Second) //返回channel类型,10秒后向channel发送当前时间
    t := <-c
    fmt.Println(t)
    tm := time.NewTimer(10 * time.Second) //NewTimer返回Timer类型
    t = <- tm.C //Timer结构中有一个channel C,10秒后把当前时间发送到C
    fmt.Println(t)

}
