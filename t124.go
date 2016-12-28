//常量

package main
import "fmt"

const PI = 3.1415926
const y = "hello"
const (
    z = false
    a = 123
)

func main(){
    const UserName,Sex = "张三","男"
    fmt.Println(PI, UserName)

}
