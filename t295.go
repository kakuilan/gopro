// 位操作
package main
import "fmt"

func main(){
    a := 0
    a |= 1 <<2 // 0000100:在bit2设置标志位
    fmt.Printf("%d %b \n",a,a)
    a |= 1 <<6 // 1000100:在bit6设置标志位
    fmt.Printf("%d %b \n",a,a)
    a = a &^ (1 << 6) // 0000100:清除bit6标志位
    fmt.Printf("%d %b \n",a,a)

}
