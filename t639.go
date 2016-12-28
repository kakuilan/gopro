//同名字段的隐藏
package main
import "fmt"

type people struct {
	name string
	sex bool
	phone string //重名字段
}

type teacher struct {
	people
	department string
	phone string //重名字段
}

func main(){
	teacher1 := teacher{
			  people{"张三", false, "100-201"},
			  "Computer Science",
			  "200-401",
			  }

	//外部字段隐藏了匿名字段同名成员
	fmt.Println(teacher1.phone)

	//通过匿名字段类型前缀访问被嵌入的同名字段
	fmt.Println(teacher1.people.phone)


}
