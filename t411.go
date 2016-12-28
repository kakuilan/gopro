// 反射Value
// Value和Type使用方法类似,包括使用Elem获取指针目标对象
package main
import (
		"fmt"
		"reflect"
		)

type User struct {
	Username string
	age int
}

type Admin struct {
	User
	title string
}

func main(){
	u := &Admin{User{"Jack", 32}, "NT"}
	v := reflect.ValueOf(u).Elem()

	fmt.Println(v.FieldByName("title").String()) //用转换方法获取字段值
	fmt.Println(v.FieldByName("age").Int()) // 直接访问嵌入字段成员
	fmt.Println(v.FieldByIndex([]int{0,1}).Int()) //用多级序号访问嵌入字段成员

}




