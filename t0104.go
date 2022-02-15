//反射修改结构体被导出的字段
package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func main() {
	t := T{23, "skiddoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	//修改字段值
	s.Field(0).SetInt(78)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now:", t)

}
