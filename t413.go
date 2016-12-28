// 反射
// 除struct,其他复合类型array,slice,map取值示例
package main
import (
		"fmt"
		"reflect"
		)

func main(){
	v := reflect.ValueOf([]int{1,2,3})

	for i,n := 0, v.Len();i<n;i++{
		fmt.Println(v.Index(i).Int())
	}

	fmt.Println("--------------------------")
	
	v = reflect.ValueOf(map[string]int{"a":1, "b":2})
	for _,k := range v.MapKeys() {
		fmt.Println(k.String(), v.MapIndex(k).Int())
	}


}


