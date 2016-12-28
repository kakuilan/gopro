//常量 iota
package main
import "fmt"

const (
    Sunday = iota
    Monday
    Tuesday
)

func main(){
    fmt.Println("Sunday=", Sunday, "Monday=", Monday, "Tuesday=", Tuesday)

    //同一行使用多个iota,它们各自增长
    const (
	U,V = iota, iota
	W,X
	Y,Z
    )
    fmt.Println(U,V,W,X,Y,Z)

    const (
	A1 = iota
	A2
	str = "Hello"
	s
	A3 = iota
	A4
    )
    fmt.Println(A1,A2,str,s,A3,A4)
}
