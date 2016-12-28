//结构体的 匿名字段
//匿名字段不过是一种语法糖,从根本上说,就是一个与成员类型同名(不包含名)的字段.被匿名嵌入的可以是任何类型,当然也包括指针
package main
import "fmt"

type User struct {
	name string
}

type Manager struct {
	User
	title string
}

func main(){
	m := Manager{
		User : User{"Tome"},	//匿名字段的显式字段名,和类型名相同
		title : "Administrator",
	}

	fmt.Println(m)
}
