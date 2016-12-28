//fallthrough语句的使用
package main
import "fmt"

func main(){
	var i int
	fmt.Println("请输入数字:")
	fmt.Scanf("%d", &i)

	switch i{
		case 0:
			fmt.Println("0")
		case 1:
			fmt.Println("1")
		case 2:
			fmt.Println("2")
			fallthrough
		case 3:
			fmt.Println("3")
		default:
			fmt.Println("4")
	}

}
