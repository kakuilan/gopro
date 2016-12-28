//函数类型

package main
import "fmt"

//MyFuncType是一个接受int类型的参数,返回bool值的函数类型
type MyFuncType func(int)bool
func IsBigThan5(n int) bool {
    return n>5
}

func Display(arr[]int, f MyFuncType){
    for _,v := range arr {
	if f(v) {
	    fmt.Println(v)
	}
    }
}

func main(){
    arr := []int{1,2,3,4,5,6,7,8,9}
    Display(arr, IsBigThan5)

}
