//// new 返回指针类型
// make 只采集slice,map,channel,并返回一个初始值的类型,而不是指针

package main
import (
    "sync"
    "bytes"
    "fmt"
)

type SyncedBuffer struct {
    lock sync.Mutex
    buffer bytes.Buffer
}

func main(){
    p := new(SyncedBuffer)
    var v SyncedBuffer
    fmt.Println(p)
    fmt.Println(v)

    var p2 *[]int = new([]int) //分配slice结构内存,很少使用
    var v2 []int = make([]int, 100) //v2指向一个新分配的有100个整数的数组

    fmt.Println(p2)
    fmt.Println(v2)




}
