// 方法
// int类型不能有方法,但是可以新建一个拥有方法的整数类型

package main
import "fmt"

type Foo int

func (self Foo) Emit(){
    fmt.Printf("%v\n", self)
}

type Emitter interface {
    Emit()
}

func main(){
    var f Foo
    f = 4
    f.Emit()

}
