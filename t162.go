// tick定时器循环
package main
import (
    "fmt"
    "time"
)

func main(){
    c := time.Tick(3 * time.Second)
    for t:= range c {
	fmt.Println(t)
    }

}
