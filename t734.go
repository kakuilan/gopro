//反射
package main
import (
    "fmt"
    "reflect"
)

type T struct {
    A int
    B string
}

func main(){
    t := T{23, "skidoo"}
    tt := reflect.TypeOf(t)
    fmt.Printf("t type:%v\n", tt)
    ttp := reflect.TypeOf(&t)
    fmt.Printf("t type:%v\n", ttp)
    //要设置t的值,需要传入t的地址,而不是t的拷贝
    //reflect.ValueOf(&t)只是一个地址的值,不是setable,通过.Elem()解引用获取t本身的reflect.Value
    s := reflect.ValueOf(&t).Elem()
    typeOfT := s.Type()
    for i:=0;i<s.NumField();i++{
        f := s.Field(i)
        fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
    }


}
