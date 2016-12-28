// 反射取stuct的tag信息
package main
import (
    "fmt"
    "reflect"
)

type Student struct {
    Name string "学生姓名"
    Age int `a:"11"b:"33"`
}

func main(){
    s := Student{}
    rt := reflect.TypeOf(s)
    fieldName,ok := rt.FieldByName("Name")
    //取tag数据
    if ok {
	fmt.Println(fieldName.Tag)
    }

    fieldAge, ok2 := rt.FieldByName("Age")
    //可以像json一样,取TAG里的数据,注意,设置时,两个之间无逗号,键名无引号
    if ok2 {
	fmt.Println(fieldAge.Tag.Get("a"))
	fmt.Println(fieldAge.Tag.Get("b"))
    }

}
