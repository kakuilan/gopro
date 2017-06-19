//append,copy
package main
import "fmt"

func printSlice(x []int){
       fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
   }

func main(){
    var numbers []int
    printSlice(numbers)

    numbers = append(numbers, 0)
    printSlice(numbers)

    numbers = append(numbers, 1)
    printSlice(numbers)

    numbers = append(numbers, 2, 3, 4)
    printSlice(numbers)

    //切片复制
    numbers1 := make([]int, len(numbers), (cap(numbers))*2)
    copy(numbers1, numbers)
    printSlice(numbers1)


}
