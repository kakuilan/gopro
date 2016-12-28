//mystruct包定义了User对象,并实现了2个对User操作的方法:SayHi(),old()
//安装 go install mystruct
package mystruct
import "fmt"

type User struct {
	Id int
	Name string
	age int
}

func (u *User) SayHi(){
	fmt.Printf("Hi, I`m %s. Nice to meet you!\n", u.Name)
}

func (u User) old() int{
	return u.age
}

