//日期
package main
import (
    "fmt"
    "time"
)

func main(){
    fmt.Println(time.Now())
    fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
    fmt.Println(time.Now().Format("07区 2006年01月02日 15时04分05秒"))

    d,err := time.Parse("01-02-2006", "07-07-2015")
    if err !=nil {
	fmt.Println(err.Error())
    }
    fmt.Println(d)
}
