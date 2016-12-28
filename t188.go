// 通过反射改变对象值
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
    s := Student{Name:"张龙", Age:20}
    //这里传的是&s 因为要修改字段的地址,否则会报错
    rv := reflect.ValueOf(&s)
    //判断是否指针类型,如果是,取指针所指向的元素的类型
    if rv.Kind() == reflect.Ptr {
	rv = rv.Elem()
    }

    //取Name字段的值
    rvField := rv.FieldByName("Name")
    fmt.Println(rvField.String())
    //设置新值
    rvField.SetString("已改名称")
    fmt.Println(s.Name)

}
