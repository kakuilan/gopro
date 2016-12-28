//map
package main
import "fmt"

func main(){
//声明一个key是字符串,值为int的字典,这种方式需要在使用之前用make初始化
var m1 map[string]int

//另一种map的声明方式
m2 := make(map[string]int)
m2["one"] = 1
m2["ten"] = 10
m2["three"] = 3

fmt.Println(m1, m2)



}
