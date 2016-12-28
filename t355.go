//结构体struct. 可以像普通字段那样访问匿名字段成员,编译器从外向内逐级查找所有层次的匿名字段,直到发现目标或出错.
package main
import "fmt"

type Resource struct {
	id int
}

type User struct {
	Resource //匿名字段
	name string
}

type Manager struct {
	User //匿名字段
	title string
}

func main(){
	var m Manager
	m.id = 1
	m.name = "Jack"
	m.title = "Administrator"

	fmt.Println(m)


}
