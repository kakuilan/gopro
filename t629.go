//结构体定义和简单操作
package main
import "fmt"

type student struct {
	id int
	name string
	sex bool
	age int
	class string
}

func main(){
	//标准方式定义
	var stu1 student
	//简写方式定义
	stu2 := student{}

	stu1.name = "李明"
	stu2.name = "张三"
	stu1.age = 18
	stu2.age = 84
	fmt.Println(stu1)
	fmt.Println(stu2)

	//结构体对象
	var stu3 *student
	stu3 = new(student)
	stu4 := new(student)
	stu3.name = "赵武"
	stu4.age = 35

	fmt.Println(stu3,stu4)

}
