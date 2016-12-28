// 反射
// 接口是否为nil,需要tab和data都为空.可使用IsNil方法判断data值
package main
import (
		"fmt"
		"reflect"
		)

func main(){
	var p *int

	var x interface{} = p
	fmt.Println(x==nil)

	v := reflect.ValueOf(p)
	fmt.Println(v.Kind(), v.IsNil())


}


