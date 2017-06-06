//自定义类型
package main
import (
    "fmt"
)

type Count int

//类型的方法
func (count *Count) Increment() {
    *count ++
}
func (count *Count) Decrement() {
    *count --
}
func (count Count) IsZero() bool {
    return count == 0
}

func main() {
    var count Count
    i := int(count)
    count.Increment()
    j := int(count)
    count.Decrement()
    k := int(count)
    fmt.Println(count, i, j, k, count.IsZero())



}



