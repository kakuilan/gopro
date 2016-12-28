//匿名字段方法
package main
import "fmt"

type People struct {
	Name string
}

type Teacher struct {
	People
	Department string
}

func (p People) SayHi() {
	fmt.Printf("Hi,I`m %s.Nice to meet you!\n", p.Name)
}

type Speaker interface {
	SayHi()
}

func main(){
	people := People{"张三"}
	teacher := Teacher{People{"李四"}, "Computer Science"}
	var is Speaker
	is = &people //为*People定义了SayHi(),自然实现该接口
	is.SayHi()
	is = &teacher //匿名字段同样也拥有SayHi()
	is.SayHi()

}
