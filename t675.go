//利用接口反射获取结构体 Type,Value,Method
package main
import (
		"fmt"
		"reflect"
		)

//定义结构体
type Student struct {
	Id	int
	Name	string
	Sex	bool
	Grade	float32
}

//为结构体定义2个方法
func (s Student) SayHi(){
	fmt.Printf("Hi,nice to meet you!\n")
}

func (s Student) MyName(){
	fmt.Printf("My name is %s, I come from China.", s.Name)
}

//反射处理函数
func StructInfo(o interface{}) {
	t := reflect.TypeOf(o)
	//判断是否为结构体类型
	if k:=t.Kind();k!=reflect.Struct {
		fmt.Printf("This value is not a struct, it`s %v", k)
		return
	}

	fmt.Println("Struct name is:", t.Name())
	fmt.Println("Fields of struct is:")

	v := reflect.ValueOf(o)
	//获取字段Type和Value信息
	for i:=0;i<t.NumField();i++{
		field := t.Field(i)
		value := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", field.Name, field.Type, value)
	}

	fmt.Println("Methods of the struct is:")
	//获取方法Name和Type信息
	for i:=0;i<t.NumMethod();i++{
		method := t.Method(i)
		fmt.Printf("%6s: %v \n", method.Name, method.Type)
	}
}

func main(){
	stu := Student{10001, "李四", false, 90.5}
	StructInfo(stu)

}
