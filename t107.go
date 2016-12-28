package main
import (
    "fmt"
    "time"
)

type Wday int
var days = []rune("日一二三四五六")
func (d Wday) String() string {
    return string(days[d])
}

func main(){
    now := time.Now()
    w := now.Weekday()
    fmt.Print("星期", Wday(w), "\n")
}
