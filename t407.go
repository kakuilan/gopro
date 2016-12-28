// 反射
// 字段标签可实现简单元数据编程,比如标记 ORM Model 属性
package main
import (
		"reflect"
		"fmt"
		)

type User struct {
	Name string `field:"username" type:"nvarchar(20)"`
	Age int `field:"age" type:"tinyint"`
}

func main(){
	var u User
	
	t := reflect.TypeOf(u)
	f,_ := t.FieldByName("Name")

	fmt.Println(f.Tag)
	fmt.Println(f.Tag.Get("field"))
	fmt.Println(f.Tag.Get("type"))



}


