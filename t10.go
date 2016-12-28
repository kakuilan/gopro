package main
import (
    "fmt"
   // "math"
)

func main(){
    var value1 complex64

    value1 = 3.2 + 12i
    value2 := 3.2 + 12i
    value3 := complex(3.2, 12)

    fmt.Println(value1, value2, value3)
    fmt.Println(real(value1), imag(value2))
}
