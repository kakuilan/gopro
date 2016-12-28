// 反射
// 对象转换为接口.
package main
import (
		"fmt"
		"reflect"
		)

type User struct {
	Username string
	age int
}

func main(){
	u := User{"Jack", 23}

	v := reflect.ValueOf(u)
	p := reflect.ValueOf(&u)

	fmt.Println(v.CanSet(), v.FieldByName("Username").CanSet())
	fmt.Println(p.CanSet(), p.Elem().FieldByName("Username").CanSet())


}





