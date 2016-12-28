// 反射
// 同样,value-interface和pointer-interface也会导致方法集存在差异
package main
import (
		"reflect"
		"fmt"
		)

type User struct {}
type Admin struct {
	User
}

func (*User) ToString(){}

func (Admin) test(){}

func main(){
	var u Admin

	methods := func(t reflect.Type){
		for i,n := 0, t.NumMethod();i<n;i++{
			m := t.Method(i)
			fmt.Println(m.Name)
		}
	}

	fmt.Println("--- value interface ---")
	methods(reflect.TypeOf(u))

	fmt.Println("--- pointer interface ---")
	methods(reflect.TypeOf(&u))
}



