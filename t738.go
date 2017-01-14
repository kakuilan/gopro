//指针
package main
import (
    "fmt"
)

func main() {
    z := 37 //z的类型为int
    pi := &z //pi的类型为*int (指向int型的指针)
    ppi := &pi //ppi的类型为**int (指向int类型指针的指针)
    fmt.Println(z, *pi, **ppi)

    **ppi++ //语意上等同于 (*(*ppi))++ 和 *(*ppi)++
    fmt.Println(z, *pi, **ppi)

}
