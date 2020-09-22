//结构体内嵌匿名成员
package main

import "fmt"

type Person struct {
	Age  int
	Name string
}

type Student struct {
	Person
}

func main() {
	per := Person{
		Age:  18,
		Name: "zhang3",
	}

	stu := Student{Person: per}
	fmt.Println("stu.Age:", stu.Age)
	fmt.Println("stu.Name:", stu.Name)
}
