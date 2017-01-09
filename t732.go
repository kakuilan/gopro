//strconv
package main
import (
    "fmt"
    "strconv"
)

func main() {
    for _,truth := range []string{"1","t","TRUE","false","F","0","5"} {
        if  b,err :=  strconv.ParseBool(truth); err !=nil {
            fmt.Printf("\n{%v}, " ,err)
        }else{
            fmt.Print(b, " ")
        }
    }

    fmt.Println()
}
