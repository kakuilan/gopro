// comma-ok类型断言
package main
import (
		"fmt"
		"strconv"
		)

type Element interface{}
type List []Element

type Person struct{
	name string
	age int
}

//打印
func (p Person) String() string{
	return "(name: "+p.name+" - age: "+strconv.Itoa(p.age)+" years)"
}

func main(){
	list := make(List, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{"Dennis", 45}

	for index,element := range list{
		switch value:= element.(type) {
			case int:
				fmt.Printf("list[%d]是整数,值为 %d\n", index,value)
			case string:
				fmt.Printf("list[%d]是字符串,值为 %s\n", index, value)
			case Person:
				fmt.Printf("list[%d]是Person类型,值为 %s\n", index, value)
			default:
				fmt.Printf("list[%d] is of a different type.")
		}
	
	}



}





