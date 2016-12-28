// 超出array最大长度
package main
import "fmt"

func main(){
    var array [100]int //索引0-99
    slice := array[0:99] //索引0-98
    slice[98] = 'a' //ok
    //slice[99] = 'b' //err,超出范围

    fmt.Println(slice)
}
