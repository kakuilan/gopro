//switch语句实现简单计算器
package main
import "fmt"

func main(){
	//操作数a,b
	var a,b int
	//算术运算符+ - * /
	var op byte

	//通过Scanf()函数,格式化从键盘输入操作数和运算符
	fmt.Println("请输入算式")
	fmt.Scanf("%d%c%d", &a,&op,&b)

	switch op{
		case '+' :
			fmt.Println(a+b)
		case '-':
			fmt.Println(a-b)
		case '*':
			fmt.Println(a*b)
		default:
			fmt.Println(a/b)
	}


}
