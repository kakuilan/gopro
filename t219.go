// slice append copy
package main
import "fmt"

func main(){
    s0 := []int{0,0}
    s1 := append(s0, 2) //追加一个元素
    s2 := append(s1, 3,5,7) //追加多个元素
    s3 := append(s2, s0...) //追加一个切片,带"..."
    fmt.Println(s1,s2,s3)

    var a = [...]int{0,1,2,3,4,5,6,7}
    var s = make([]int, 6)
    n1 := copy(s, a[0:]) //n1==6, s== []int{0,1,2,3,4,5}
    fmt.Println(n1,s)
    n2 := copy(s, s[2:]) //n2==4, s== []int{2,3,4,5,4,5}
    fmt.Println(n2,s)

}
