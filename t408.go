// 反射
// 可从基本类型获取所对应复合类型
package main
import (
		"fmt"
		"reflect"
		)

var (
		Int = reflect.TypeOf(0)
		String = reflect.TypeOf("")
		)

func main(){
	c := reflect.ChanOf(reflect.SendDir, String)
	fmt.Println(c)

	m := reflect.MapOf(String, Int)
	fmt.Println(m)

	s := reflect.SliceOf(Int)
	fmt.Println(s)

	t := struct{Name string}{}
	p := reflect.PtrTo(reflect.TypeOf(t))
	fmt.Println(p)



}




