//反射.
//以struct为例,可获取其全部成员字段信息,包括非导出和匿名字段
package main
import (
		"fmt"
		"reflect"
		)


type User struct {
	Username string
}

type Admin struct {
	User
	title string
}

func main(){
	var u Admin
	t := reflect.TypeOf(u)

	for i,n :=0,t.NumField();i<n;i++{
		f := t.Field(i)
		fmt.Println(f.Name, f.Type)
	}


}





