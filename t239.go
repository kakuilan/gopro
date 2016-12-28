// even 包测试
package main
import (
    "even" //导入本地even包
    "fmt"
)

func main(){
    i := 5
    fmt.Println("i的奇偶:", i, even.Even(i))

}
