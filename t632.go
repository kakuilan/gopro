//相同类型的结构体直接可直接使用"="操作符进行赋值
package main
import "fmt"

type user struct {
	id int
	name string
}

func main(){
	user1 := user{100,"张三"}
	fmt.Println(user1)

	var user2 *user = new(user)
	* user2 = user1
	fmt.Println(user2)
}
