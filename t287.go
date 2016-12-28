// 
package main
import "fmt"

func main(){
    u := uint32(32)
    i := int32(1)
    fmt.Println(&u, &i) 

    //以下将报错
    p := &i // p的类型是 *int32
    p = &u // &u的类型是 *uint32, 和p的类型不同,不能赋值
    p = (*int32)(&u) //这种类型转换语法也是无效的
    fmt.Println(p)


}
