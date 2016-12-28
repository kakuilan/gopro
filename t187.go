// 反射获取值
package main
import (
    "fmt"
    "reflect"
)

type Student struct {
    Name string
    Age int
}

func main(){
    s := Student{Name:"李四",Age:29}
    rv := reflect.ValueOf(s)
    //判断是否指针类型,如果是,取指针所指向的元素的类型
    if rv.Kind() == reflect.Ptr {
	rv = rv.Elem()
    }
    rvField := rv.FieldByName("Name") //取Name字段的值
    fmt.Println(rvField.String())

}
