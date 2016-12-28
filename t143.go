//方法继承的问题
package main
import "fmt"

type Fruit struct {
}

func (this *Fruit) DisplayName(){
    fmt.Println(this.GetName())
}

func (this *Fruit) GetName() string {
    return "水果"
}

type Apple struct {
    Fruit
}

func (this *Apple) GetName() string{
    return "苹果"
}

func main(){
    fruit := Fruit{}
    fruit.DisplayName()
    apple := Apple{}
    apple.DisplayName()
}
