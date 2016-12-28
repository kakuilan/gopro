package main
import (
    "fmt"
    "math"
    "reflect"
)

//通过反射包为一个结构体变量赋值
type X struct {
    Y byte
    z complex128
}

func main(){
    var x X
    fmt.Println(x)

    v := reflect.ValueOf(&x)
    e := v.Elem()
    a := math.Pi
    fmt.Println(v,e,a)
    return
    //e.Field(0).SetUnit('e')

    p := math.Pi
    c := complex(math.Cos(p), math.Sin(p))
    e.Field(1).SetComplex(c)
    fmt.Println("最美的公式:")
    fmt.Println("%c^ip+1=%g\n", x.Y, 1+real(x.z))
    fmt.Println("可惜虚部有点精度误差:", imag(x.z))
}
