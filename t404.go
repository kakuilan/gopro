//反射.
//如果是指针,应该先使用Elem方法获取目标类型,指针本身是没有字段成员的
package main
import (
		"fmt"
		"reflect"
		)


type User struct {
	Username string
}

type Admin struct {
	User
	title string
}

func main(){
	u := new(Admin) //返回指针
	t := reflect.TypeOf(u)
	
	//这里判断释放指针
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	


	for i,n :=0,t.NumField();i<n;i++{
		f := t.Field(i)
		fmt.Println(f.Name, f.Type)
	}


}





