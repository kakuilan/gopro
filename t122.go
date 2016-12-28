//切片
package main
import (
    "fmt"
    "reflect"
)

func main(){
    //定义一个数组
    p := [...]int{2,3,5,7,11,13}
    
    //定义切片,包含3,5两个元素
    s1 := p[1:3]
    fmt.Println(s1)

    //用反射得到变量的类型P是数组类型[6]int
    fmt.Println(reflect.TypeOf(p))

    //s1是切片类型[]int
    fmt.Println(reflect.TypeOf(s1))

    //ChangeArrayValue函数将第一个值改为100
    ChangeArrayValue(p)
    //数组p的值并没有改变,因为数组是值类型
    fmt.Println(p)

    //ChangeSliceValue将切片的第一个值改为100
    ChangeSliceValue(s1)
    //切片s1的值被改变,因为切片是引用类型
    fmt.Println(s1)
    //切片是引用的数组p第一个和第二个元素,所以数组p的值被改变
    fmt.Println(p)

}

func ChangeArrayValue(arr [6]int) {
    arr[0] = 100
}

func ChangeSliceValue(slice []int) {
    slice[0] = 100
}
