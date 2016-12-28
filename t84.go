package main
import (
    "fmt"
    "reflect"
)

type X struct {
    y byte
    z complex128
}

func (x *X) String() string {
    return fmt.Sprintf("%v", x)
}

func main(){
    var x X
    fmt.Println(x)

    v := reflect.TypeOf(x)
    fmt.Println("type:", v)
    fmt.Println("Align:", v.Align())
    fmt.Println("FieldAlign:", v.FieldAlign())

    for i := 0;i<v.NumMethod();i++{
	fmt.Println("Method ", i, v.Method(i).Name)
    }

    fmt.Println("Name:", v.Name())
    fmt.Println("PkgPath:", v.PkgPath())
    fmt.Println("Kind:", v.Kind())

    for i:=0;i<v.NumField();i++{
	fmt.Println("Field ", i, v.Field(i).Name)
    }
    fmt.Println("Size:", v.Size())


}
