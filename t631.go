//结构体变量、对象的初始化
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
	//全部字段初始化
	stu1 := student{13001,"李明",false,19,"网络01"} //结构体变量
	stu2 := &student{13002,"张欣",true,18,"网络02"} //结构体对象

	//部分初始化
	stu3 := student{name:"王丽", age:19} //结构体变量
	stu4 := &student{name:"赵传",id:13003} //结构体对象

	fmt.Println(stu1)
	fmt.Println(stu2)
	fmt.Println(stu3)
	fmt.Println(stu4)
}

