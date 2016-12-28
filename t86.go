package main
import (
    "fmt"
    "reflect"
)

func main(){
    var x float64 = 3.14
    fmt.Println("x:", x)

    v := reflect.ValueOf(&x)
    fmt.Println("v.CanSet:", v.CanSet())
    i := v.Elem()
    fmt.Println("i.CanSet:", i.CanSet())
    i.SetFloat(3.1415)
    fmt.Println("i:", i.Float())
    fmt.Println("x:", x)
}
