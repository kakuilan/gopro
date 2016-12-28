//结构体不能同时嵌入某一类型和其指针类型,因为它们名字相同.
package main
import "fmt"

type Resource struct {
	id int
}

type User struct {
	*Resource
	//Resource //Error: duplicate file Resource
	name string
}

func main(){
	u := User{
		&Resource{1},
		"Administrator",
	}

	println(u.id)
	println(u.Resource.id)
	fmt.Println(u)

}

