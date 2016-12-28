//通用格式化输出
package main
import "fmt"

func main(){
	var id int = 100
	var name string = "李明"
	var grade float32 = 91.5
	fmt.Println("通用格式化输出：")
	fmt.Printf("%v %v %v\n", id,name,grade)
	fmt.Printf("%#v %#v %#v\n", id,name,grade)
	fmt.Printf("%T %T %T\n", id,name,grade)
	fmt.Printf("60 %%\n")

	var yes,no = true,false
	fmt.Println("布尔型数据输出：")
	fmt.Printf("%t %t\n", yes,no)
}
