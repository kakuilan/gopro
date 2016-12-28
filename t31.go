package main
import "fmt"
//不定参数
func myFunc(args ...int){
    for _, arg := range args{
	fmt.Println(arg)
    }
}

func main(){
    myFunc(12,4,45,567,77)
}
