//使用无条件表达式switch语句打印成绩等级
package main
import "fmt"

func main(){
	var grade int
	fmt.Println("请输入分数:")
	fmt.Scanf("%d\n", &grade)

	switch{
		case (grade>=90) && (grade<=100) :
			fmt.Println("等级A")
		case (grade>=80) && (grade<90):
			fmt.Println("等级B")
		case (grade>=70) && (grade<80):
			fmt.Println("等级C")
		default:
			fmt.Println("等级D")
	
	
	}


}
