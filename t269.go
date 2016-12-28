// 编写一个Quine，打印自己的程序
package main
import "fmt"
func main(){
    fmt.Printf("%s%c%s%c\n", q, 0x60, q, 0x60)
}

var q = `
package main
import "fmt"
func main(){
    fmt.Printf("%s%c%s%c\n", q, 0x60, q, 0x60)
}
var q = `

