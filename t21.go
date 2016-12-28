package main
import "fmt"

func main(){
    mySlice1 := make([]int, 5)
    mySlice2 := make([]int, 5, 10)
    mySlice3 := []int{1,2,3,4,5}

    fmt.Println(mySlice1, mySlice2, mySlice3)
    fmt.Println("len(mySlice2):", len(mySlice2))
    fmt.Println("cap(mySlice2):", cap(mySlice2))

    mySlice1 = append(mySlice1, 1,2,3)
    var mySlice4 = []int{8,9,10}
    var mySlice5 = append(mySlice3,mySlice4...) //mySlice4后面加了3个点,即一个省略号,如果没有这个省略号,会编译错误,原因看书 
    fmt.Println(mySlice1, mySlice5)

    for i:=0;i<len(mySlice1);i++{
	fmt.Println("mySlice1[",i,"]=", mySlice1[i])
    }
    for i,v := range mySlice3 {
	fmt.Println("mySlice3[",i,"]=", v)
    }


}
