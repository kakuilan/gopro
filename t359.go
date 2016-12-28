// 方法
package main
import "fmt"

type Queue struct {
	elements []interface{}
}

func NewQueue() *Queue{	//创建对象实例
	return &Queue{make([]interface{}, 10)}
}

func (*Queue) Push(e interface{}) error { //省略receiver参数名
	panic("not implemented")
}

//错误,不支持方法重载 
// Error: method redeclared: Queue.Push
//func (Queue) Push(e int) error {
//	panic("not implemneted")
//}

func (self *Queue) length() int {
	return len(self.elements)
}

func main(){
	a := NewQueue()
	fmt.Println(a)

	var b int
	b = a.length()
	fmt.Println(b)

}
