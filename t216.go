// array
package main
import "fmt"

func main(){
    var arr [10]int
    arr[0] = 23
    arr[1] = 33
    s := "abc"
    fmt.Println(arr,s)

    a2 := [...]int{1,2,3,4}
    //多维数组
    a3 := [3][2] int {[2]int{1,2}, [2]int{3,4}, [2]int{5,6}}
    var a4 [4][2]int
    fmt.Println(a2, a3, a4)
}
