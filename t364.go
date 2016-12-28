// instance.method(args...) --> <type>.func(instance, args...)
// 前者称为method value, 后者method expression
// 区别在于method value绑定实例,而method expression则必须显式传参
package main
import "fmt"

type User struct {
	id int
	name string
}

func (self *User) Test(){
	fmt.Printf("%p, %v\n", self, self)
}

func main(){
	u := User{3, "Hamu"}
	u.Test() //方法

	mValue := u.Test
	mValue()	//隐式传递 receiver

	mExpression := (*User).Test
	mExpression(&u) //显式传递 receiver




}


