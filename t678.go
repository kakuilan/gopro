//利用反射修改原对象Value值
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Id int
	Name string
	Sex bool
	Grade float32
}

func SetValue(o interface{}) {
	v := reflect.ValueOf(o)
	if v.Kind() != reflect.Ptr || !v.Elem().CanSet() {
		fmt.Println("Cannot set!")
		return
	}else{
		v = v.Elem()
	}

	for i:=0;i<v.NumField();i++{
		switch v.Field(i).Kind() {
			case reflect.Int:
				v.Field(i).SetInt(10002)
			case reflect.String:
				v.Field(i).SetString("张翰")
			case reflect.Bool:
				v.Field(i).SetBool(true)
			case reflect.Float32:
				v.Field(i).SetFloat(95.5)
		}
	}
}

func main(){
	stu := Student{100, "李明", false, 88.1}
	fmt.Println("修改前：", stu)
	SetValue(&stu)
	fmt.Println(stu)
}
