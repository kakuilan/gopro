package main
import "fmt"

func main(){
    var iarray1 [5]int32
    var iarray2 [5]int32 = [5]int32{1,2,3,4,5}
    iarray3 := [5]int32{1,2,3,4,5}
    iarray4 := [5]int32{6,7,8,9,10}
    iarray5 := [...]int32{11,12,13,14,15}
    iarray6 := [4][4]int32{{1},{1,2},{1,2,3}}

    fmt.Println(iarray1)
    fmt.Println(iarray2)
    fmt.Println(iarray3)
    fmt.Println(iarray4)
    fmt.Println(iarray5)
    fmt.Println(iarray6)


}


