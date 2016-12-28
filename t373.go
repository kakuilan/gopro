// 只有tab和data都为nil时,接口才等于nil
package main
import "fmt"
import "unsafe"
import "reflect"


func main(){

	
var a interface{} = nil //tab=nil, data=nil
var b interface{} = (*int)(nil) //tab包含*int类型信息, data=nil

type iface struct {
	itab, data uintptr
}

ia := *(*iface)(unsafe.Pointer(&a))
ib := *(*iface)(unsafe.Pointer(&b))
	
	fmt.Println(a==nil, ia)
	fmt.Println(b==nil, ib, reflect.ValueOf(b).IsNil())

}


