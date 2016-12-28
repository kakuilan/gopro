// 接口转换.利用类型推断,可判断接口对象是否某个具体的接口或类型
package main
import "fmt"

type User struct {
	id int
	name string
}

func (self *User) String() string {
	return fmt.Sprintf("%d, %s", self.id, self.name)
}

func main(){
	var o interface{} = &User{1, "Tom"}

	if i,ok := o.(fmt.Stringer);ok { //判断是不是fmt包的Stringer接口(实现了String()方法即是)
		fmt.Println(i)
	}

	u := o.(*User)
	// u := o.(User) // panic: interface is *main.User, not main.User
	fmt.Println(u)


}




