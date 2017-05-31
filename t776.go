//可变参数
package main
import (
    "fmt"
)

func MinimumIntl(first int,rest ...int) int{
    for _,x := range rest {
        if x < first {
            first = x
        }
    }

    return first
}

func main() {
    fmt.Println(MinimumIntl(5,3), MinimumIntl(7,3,-2,4,0,-8,-5))

}
