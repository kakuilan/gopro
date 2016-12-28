// 反射调用函数
package main
import (
    "fmt"
    "reflect"
)

type Student struct {
    Name string
    Age int
}

func (this *Student) PrintName(){
    fmt.Println(this.Name)
}

func (this *Student) GetAge() int {
    return this.Age
}

func main(){
    s := Student{Name:"赵虎", Age:21}
    rt := reflect.TypeOf(&s)
    rv := reflect.ValueOf(&s)

    fmt.Println("TypeOf调用函数")
    rtm,ok := rt.MethodByName("PrintName")
    if ok {
	var parm []reflect.Value
	//函数默认第一个参数是结构体本身,即 *Student
	parm = append(parm, rv)
	rtm.Func.Call(parm)
    }

    fmt.Println("ValueOf调用函数")
    rvm := rv.MethodByName("GetAge")
    //用valueof调用函数不需要把Struct本身作为参数传递过去
    ret := rvm.Call(nil)
    //显示返回值
    fmt.Println("返回值")
    ShowSlice(ret)
}

func ShowSlice(s []reflect.Value) {
    if s != nil && len(s) >0 {
	for _,v := range s {
	    fmt.Println(v.Interface())
	}
    }
}
