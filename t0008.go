//定义一个学生结构体
package main

import "fmt"

type Student struct {
	Age  int
	Name string
}

func main() {
	stu := Student{
		Age:  18,
		Name: "Li4",
	}

	fmt.Println(stu)

	//在赋值的时候,字段名可以忽略
	fmt.Println(Student{20, "Zhang3"})
	return
}
