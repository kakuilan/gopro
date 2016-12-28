//匿名字段的访问
package main
import "fmt"

type people struct {
	name string
	sex bool
}

type teacher struct {
	people //匿名字段
	department string
}

type departmentHead struct {
	teacher //有2层嵌入的匿名字段
	college string
}

func main(){
	head := departmentHead{}
	head.name = "李四"
	head.department = "Computer Science"
	head.college = "Yale University"
	fmt.Println(head)
}
