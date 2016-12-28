// 从map中取回的是一个value临时复制品,对其成员的修改是没有任何意义的.
package main
import "fmt"

func main(){
	type user struct {name string}

	m := map[int] user {	//当map因扩张而重新哈希时,各键值项存储位置都会发生改变.因此,map
		1 : {"user1"},		//被设计成not addressable. 类似m[1].name这种期望透过原value
	}						//指针修改成员的行为自然会被禁止

	//下面错误
	//m[1].name = "Tom" //

	//正确做法是完整替换value或使用指针
	u := m[1]
	u.name = "Tom"
	m[1] = u //替换value
	fmt.Println(m)

	//使用指针
	m2 := map[int] *user{
		1 : &user{"user2"},
	}
	fmt.Println(m2)
	m2[1].name = "Jack" //返回的是指针复制品.透过指针修改原对象是允许的
	fmt.Println(m2)

}


