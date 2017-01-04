// %T %v格式
package main
import (
    "fmt"
    "math"
)

func main(){
    type polar struct {a,b float32}
    p := polar{-83.40, 71.60}

    fmt.Printf("|%T|%v|%#v|\n", p, p, p)
    fmt.Printf("|%T|%v|%t|\n", false, false, false)
    fmt.Printf("|%T|%v|%f|\n", 7607, 7607, 7607)
    fmt.Printf("|%T|%v|%f|\n", math.E, math.E, math.E)
    fmt.Printf("|%T|%v|%f|\n", 5+7i, 5+7i, 5+7i)
    s := "Relativity"
    fmt.Printf("|%T|\"%v\"|\"|%q|\n", s, s, s, s)


}
