//多返回值

package main
import "fmt"

func divide(a,b int) (int,int) {
    quotient := a /b
    remainder := a % b
    return quotient,remainder
}

func main(){
    q,r := divide(5,3)
    fmt.Println(q, ",", r)
}
