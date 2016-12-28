//外部Struct的导入
package main
import (
		"fmt"
		"mystruct"
		)

func main(){
	user := new(mystruct.User)
	user.Id = 100
	user.Name = "张三"
	user.SayHi()
	fmt.Println(user)


}
