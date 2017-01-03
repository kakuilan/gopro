//打印布尔值
package main
import (
    "fmt"
)

func IntForBool(b bool) int {
    if b {
        return  1
    }else{

        return  0
    }

}

func main(){
    fmt.Printf("%t %t\n", true, false)

    fmt.Printf("%d %d\n", IntForBool(true), IntForBool(false))

}
