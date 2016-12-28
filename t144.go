//继承
package main
import "fmt"

//定义一个函数类型
type FruitName func() string

type Fruit struct {
    GetFruitName FruitName
}

func (this *Fruit) DisplayName(){
    fmt.Println(this.GetFruitName())
}

func (this *Fruit) GetName() string {
    return "水果"
}

func NewFruit() *Fruit {
    f := new(Fruit)
    f.GetFruitName = f.GetName
    return f
}

type Apple struct {
    Fruit
}

func (this *Apple) GetName() string {
    return "苹果"
}

func NewApple() *Apple {
    a := new(Apple)
    a.GetFruitName = a.GetName
    return a
}

func main(){
    fruit := NewFruit()
    fruit.DisplayName()
    apple := NewApple()
    apple.DisplayName()

}
