package main
import (
    "fmt"
    "math"
)

func IsEqual(f1, f2, p float32) bool {
    return math.Fdim(f1, f2) < p
}

func main(){
    var fvalue1 float32
    fvalue1 = 12
    fvalue2 := float32(12.0)
    var res bool
    res = IsEqual(fvalue1, fvalue2, 0.0)
    fmt.Println(fvalue1, fvalue2, res)
}
