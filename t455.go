// comma-ok类型断言
package main
import (
		"fmt"
		"strconv"
		)

type Element interface{}
type List []Element

type Person struct {
	name string
	age int
}

//定义String方法,实现了fmt.Stringer
func (p Person) String() string {
	return "(name: " + p.name + " - age: "+strconv.Itoa(p.age) +" years)"
}

func main(){
	list := make(List, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{"Dennis", 59}

	for index,element := range list {
		if value,ok := element.(int); ok{
			fmt.Printf("list[%d]是一个整数,值为 %d\n", index, value)
		}else if value,ok := element.(string);ok{
			fmt.Printf("list[%d]是一个字符串,值为 %s\n", index, value)
		}else if value,ok := element.(Person);ok{
			fmt.Printf("list[%d]是一个Person类型,值为 %s\n", index, value)
		}else{
			fmt.Printf("list[%d] is of a different type",index)
		}
	
	}

}




