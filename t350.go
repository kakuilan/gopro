// struct顺序初始化必须包含全部字段,否则会出错.
package main
import "fmt"

type User struct {
	name string
	age int
}


func main(){

	u1 := User{"Tom", 22}
	//u2 := User{"Kev"} //Error too few values in struct initializer

	fmt.Println(u1)

}
