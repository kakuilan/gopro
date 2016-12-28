// 新类型不是原类型的别名,除拥有相同数据存储结构外,它们之间没有任何关系,不会持有原类型任何信息.除非模板类型是未命名类型,否则必须显式转换
package main
import "fmt"

type bigint int64
func main(){
    x := 1234
    var b bigint = bigint(x) //必须显式转换,除非是常量
    var b2 int64 = int64(b)

    var s myslice = []int{1,2,3} //未命名类型,隐式转换
    var s2 []int = s

    fmt.Println(b,b2,s,s2)


}
