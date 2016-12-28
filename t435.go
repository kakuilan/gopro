// GO的if语句里面允许声明一个变量,该变量的作用域只能在该条件逻辑块内
package main
import "fmt"

func main(){

if x := 1; x>10 {
	fmt.Println("x is greater than 10")
}else{
	fmt.Println("x is less than 10")
}

}

