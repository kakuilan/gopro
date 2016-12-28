//反射
// 方法Implements判断是否实现了某个具体接口,AssignableTo,ConvertibleTo用于赋值和转换判断
package main
import (
		"fmt"
		"reflect"
		)

type Data struct {}

func (*Data) String() string {
	return ""
}

func main(){
	var d *Data
	t := reflect.TypeOf(d)

	//没法直接获取接口类型,好在接口本身是个struct
	//创建一个空指针对象,这样传递给TypeOf转换成 interface{}时就有类型信息了
	it := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

	//为啥不是 t.Implements(fmt.Stringer),完全可以由编译器生成
	fmt.Println(t.Implements(it))


}



