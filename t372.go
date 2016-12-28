// 接口转型返回临时对象,只有使用指针才能修改其状态
package main
import "fmt"

type User struct {
	id int
	name string
}

func main(){
	u := User{1, "Tome"}
	var vi, pi interface{} = u, &u
	
	// vi.(User).name = "Jack" // Error: cannot assign to vi.(User).name
	pi.(*User).name = "Jack"

	fmt.Printf("vi: %v\n", vi.(User))
	fmt.Printf("pi: %v\n", pi.(*User))
	fmt.Println(vi, pi)
}
