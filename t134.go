//多返回值的命名

package main
import "fmt"

func divide(a,b int) (quotient, remainder int) {
    quotient = a /b
    remainder = a % b
    return
}

func main(){
    q,r := divide(5,3)
    fmt.Println(q, ",", r)
}
