// 反射Method
// 如果改用CallSlice,只需将变参打包成slice即可
package main
import (
		"fmt"
		"reflect"
		)

type Data struct {}

func (*Data) Test(x, y int) (int, int) {
	return x+100, y+100
}

func (*Data) Sum(s string, x ...int) string {
	c := 0
	for _,n := range x{
		c +=n
	}
	return fmt.Sprintf(s, c)
}

func info(m reflect.Method) {
	t := m.Type
	fmt.Println(m.Name)

	for i,n := 0,t.NumIn();i <n;i++{
		fmt.Printf(" in[%d] %v\n", i, t.In(i))	
	}

	for i,n := 0,t.NumOut();i<n;i++{
		fmt.Printf(" out[%d] %v\n", i, t.Out(i))
	}
}

func main(){
	d := new(Data)
	v := reflect.ValueOf(d)

	m := v.MethodByName("Sum")
	in := []reflect.Value{
		reflect.ValueOf("result = %d"),
		reflect.ValueOf([]int{1,2}), //将变参打包成slice
	}
	   
	out := m.CallSlice(in)
	fmt.Println(out)
	for _,v = range out {
		fmt.Println(v.Interface())
	}

}




