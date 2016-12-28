package main
import "fmt"

func main(){
    str := "Hello,世界"
    for i, ch := range str {
	fmt.Println(i, ch);//ch的类型为rune
    }
}
