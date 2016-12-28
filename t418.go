// 反射
// 复合类型修改示例
package main
import (
		"fmt"
		"reflect"
		)

func main(){
	s := make([]int, 0, 10)
	v := reflect.ValueOf(&s).Elem()

	v.SetLen(2)
	v.Index(0).SetInt(100)
	v.Index(1).SetInt(200)

	fmt.Println(v.Interface(), s)

	v2 := reflect.Append(v, reflect.ValueOf(300))
	v2 = reflect.AppendSlice(v2, reflect.ValueOf([]int{400, 500}))

	fmt.Println(v2.Interface())

	fmt.Println("-------------------------")

	m := map[string]int{"a":1}
	v = reflect.ValueOf(&m).Elem()

	v.SetMapIndex(reflect.ValueOf("a"), reflect.ValueOf(100)) //update
	v.SetMapIndex(reflect.ValueOf("b"), reflect.ValueOf(200)) //add

	fmt.Println(v.Interface(), m)

}


