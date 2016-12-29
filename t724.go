//字符串格式化
package main
import (
    "fmt"
)

func main(){
    type polar struct {radius,o float64}
    p := polar {8.32, .49}
    fmt.Print(-18.5, 17, "Elephant", -8+.7i, 0x3c7, '\u03c7', "a", "b", p)
    fmt.Println()
    fmt.Println(-18.5, 17, "Elephant", -8+.7i, 0x3c7, '\u03c7', "a", "b", p)


}
