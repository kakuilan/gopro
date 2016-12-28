// 面向对象
package main
import "fmt"

type User struct {
	id int
	name string
}

type Manager struct {
	User
	title string
}

func main(){
	m := Manager{User{1, "Tom"}, "Administrator"}

	//var u User = m //Error: cannot use m (type Manager) as type User in assignment
					//没有继承,自然也不会有多态
	var u User = m.User //同类型拷贝
	fmt.Println(m, u)


}


