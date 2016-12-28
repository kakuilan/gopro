//时间包
package main
import (
    "fmt"
    "time"
)

func main(){
    now := time.Now()
    fmt.Println(now)
    fmt.Println(now.Format(time.RFC3339))
    fmt.Println(now.Format("01月02日03时04分05秒06年07区"))
}
