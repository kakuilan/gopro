//利用反射动态调用原对象方法
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

func (s Student) SayHi(name string) {
	fmt.Printf("Hi %s, my name is %s.\n", name, s.Name)
}

func main(){
	stu := Student{10001, "liMing", false, 89.2}
	v := reflect.ValueOf(stu)
	mv := v.MethodByName("SayHi")
	args := []reflect.Value{reflect.ValueOf("张三")}
	mv.Call(args)
}
