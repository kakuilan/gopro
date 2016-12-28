package main
import "fmt"

func main(){
    oldSlice := []int{1,2,3,4,5}
    newSlice1 := oldSlice[:3] //基于oldSlice的前3个元素构建新数组切片
    newSlice2 := oldSlice[1:]

    fmt.Println(oldSlice, newSlice1, newSlice2)

    slice1 := []int{1,2,3,4,5}
    slice2 := []int{5,4,3}
    slice3 := []int{5,4,4}
    copy(slice2, slice1)
    fmt.Println(slice1, slice2)
    copy(slice1, slice3)
    fmt.Println(slice1, slice3)

}
