package main
import "fmt"

func main(){
    iarray7 := [5]string{"aaa", "bb", "可以了", "叫我说什么好", "()"}
    fmt.Println(iarray7)
    for i:= range iarray7 {
	fmt.Println(iarray7[i])
    }

}
