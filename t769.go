//自定义类型
package main
import "fmt"

type StringSlice []string
func (s StringSlice) String() string {
    return "test"
}

func main() {
    fancy := StringSlice{"Lithin", "Sodiun", "Postass", "Rubiit"}
    fmt.Println(fancy)
    plain := []string(fancy)
    fmt.Println(plain)
}
