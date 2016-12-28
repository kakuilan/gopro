//数组

package main
import "fmt"

func main(){
    var arr_1 [2]int
    fmt.Println(arr_1)

    arr_2 := [2]int {}
    fmt.Println(arr_2)

    arr3 := [2]int{1,2}
    arr3_1 := [2]int{0:1, 1:2}
    arr3_2 := [2]int{1:2, 0:1}
    fmt.Println(arr3, arr3_1, arr3_2)

    //不指定数组长度
    arr4 := [...]int{1,2,3}
    fmt.Println(arr4)

    arr5 := [...]int{3:9}
    fmt.Println(arr5)




}
