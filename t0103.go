//反射结构体
package main

import (
	"fmt"
	"reflect"
)

type NotknownType struct {
	s1, s2, s3 string
}

func (n NotknownType) String() string {
	return n.s1 + " - " + n.s2 + " - " + n.s3
}

//variable to investigate
var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

func main() {
	value := reflect.ValueOf(secret)
	typ := reflect.TypeOf(secret)
	knd := value.Kind()

	fmt.Println(typ)
	fmt.Println(knd)

	//遍历结构体字段
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
	}

	//调用结构体方法
	results := value.Method(0).Call(nil)
	fmt.Println(results)

}
