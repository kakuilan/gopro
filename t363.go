//通过匿名字段,可获得和继承类似的复用能力.依据编译器查找次序,只需在外层定义同名方法,就可以实现override
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

func (self *User) ToString() string {
	return fmt.Sprintf("User: %p, %v", self, self)
}


//这里重写User的ToString方法
func (self *Manager) ToString() string {
	return fmt.Sprintf("Manager: %p, %v", self, self)
}

func main(){
	m := Manager{User{2, "Jari"}, "Administrator"}
	fmt.Println(m.ToString())
	fmt.Println(m.User.ToString())
}




