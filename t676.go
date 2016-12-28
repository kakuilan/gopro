//结构体的嵌入字段反射 reflect FieldByIndex()
package main
import (
		"fmt"
		"reflect"
		)

type Student struct {
	Id	int
	Name string
	Sex bool
	Grade float32
}

type Monitor struct {
	Student
	As string
}

func main(){
	stu := Monitor{Student{10001,"李四",false,90.5}, "班长"}
	t := reflect.TypeOf(stu)
	v := reflect.ValueOf(stu)
	for i:=0;i<t.NumField();i++{
		if t.Field(i).Anonymous { //Anonymous方法判断是否匿名字段
			for j:=0;j<v.Field(i).NumField();j++{
				fmt.Println("Embedded fields:", v.FieldByIndex([]int{i,j}).Interface())
			}	
		}else{
			fmt.Println("Normal fields:", v.Field(i).Interface())
		}
	}

}
