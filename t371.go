// 接口执行机制.
// 接口对象由接口表(interface table)指针和数据指针组成
// 数据指针持有的是目标对象的只读复制品,复制完整对象或指针
package main
import "fmt"

type User struct {
	id int
	name string
}

func main(){
	u := User{1, "Tom"}
	var i interface{} = u

	u.id = 2
	u.name = "Jack"

	fmt.Printf("u:%v\n", u)
	fmt.Printf("i:%v\n", i.(User))
	fmt.Printf("i:%v\n", i)


}
