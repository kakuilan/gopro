package main
import (
    "fmt"
    "math"
    "strconv"
)

func main(){
    fmt.Println(strconv.FormatUint(1, 2))
    fmt.Println(strconv.FormatUint(math.Float64bits(1e4),2))

    var i = 0+1i
    fmt.Println(i*i)

    var a = complex(3.14, -1)
    var b = imag(a)
    var c = real(a)
    fmt.Println(a, b, c) 
}
