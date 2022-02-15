//通过反射修改值
package main

import (
	"fmt"
	"reflect"
)

func main() {
	//检查是否能设置
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("settability of v:", v.CanSet())

	//必须使用指针
	v = reflect.ValueOf(&x)
	fmt.Println("type of v:", v.Type())
	fmt.Println("settability of v:", v.CanSet())

	v = v.Elem()
	fmt.Println("The Elem of v is:", v)
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(3.14159)
	fmt.Println(v.Interface())
	fmt.Println(v)

}
