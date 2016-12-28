// time时间
package main
import (
    "fmt"
    "time"
)

func main(){
    t := time.Now()
    t2 := t.Add(24 * time.Hour) //当前时间加24小时,即明天的此时
    d := t2.Sub(t) //t2-t,相差24小时
    fmt.Println(t)
    fmt.Println(t2)
    fmt.Println(d)
    if t.Before(t2) {
	fmt.Println("t < t2")
    }

    if t.Equal(t) {//判断两个时间是否相等
	fmt.Println("t=t")
    }

}
