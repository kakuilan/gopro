// 反射
// 除具体的Int,String等转换方法,还可返回interface{}.只是非导出字段无法使用,需用CanInterface判断下
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

	//报错
	//fmt.Println(v.FieldByName("Username").Interface())
	//fmt.Println(v.FieldByName("age").Interface())

	f := v.FieldByName("age")
	if f.CanInterface(){
		fmt.Println(f.Interface())
	}else{
		fmt.Println(f.Int())
	}


}



