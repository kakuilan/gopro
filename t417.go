// 反射
// 非导出字段无法直接修改,可改用指针操作.
package main
import (
		"fmt"
		"reflect"
		"unsafe"
		)

type User struct {
	Username string
	age int
}

func main(){
	u := User{"Jack", 23}
	p := reflect.ValueOf(&u).Elem()

	p.FieldByName("Username").SetString("Tom")

	f := p.FieldByName("age")
	fmt.Println(f.CanSet())

	//判断是否能获取地址
	if f.CanAddr() {
		age := (*int) (unsafe.Pointer(f.UnsafeAddr()))
		// age := (*int)(unsafe.Ponter(f.Addr().Pointer())) //等同
		*age = 88
	}

	//注意p是Value类型,需要还原成接口才能转型
	fmt.Println(u, p.Interface().(User))

}


