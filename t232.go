// 整数顺序, 小->大
package main
import "fmt"

func order(a,b int) (int,int) {
    if a>b {
	return b,a
    }
    return a,b
}

func main(){
    a,b := 6,2
    fmt.Println(a,b)
    a,b = order(a,b)
    fmt.Println(a,b)
}
