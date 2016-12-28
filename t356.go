//结构体struct.匿名字段.
//外层同名字段会遮蔽嵌入字段成员,相同层次的同名字段也会让编译器无所适从.解决方案是使用显式字段名.
package main
import "fmt"

type Resource struct {
	id int
	name string
}

type Classify struct {
	id int
}

type User struct {
	Resource	//Resource.id与Classify.id处于同一层次
	Classify
	name string	//遮蔽Resource.name
}


func main(){
	u := User{
		Resource{1, "people"},
		Classify{100},
		"Jack",	   
	}

	println(u.name)	//User.name: Jack
	println(u.Resource.name) //people

	//println(u.id)	//Error: ambiguous selector u.id
	println(u.Classify.id) //100

	fmt.Println(u)



}
