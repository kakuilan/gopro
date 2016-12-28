//结构体支持 "=="、"!="相等操作符,可用作maop键类型
package main
import "fmt"

type User struct {
	id int
	name string
}

func main(){
	m := map[User]int{
		User{1, "Tom"} : 100,
	}

	fmt.Println(m)



}
