//从键盘输入3个整数,求平均值,保留2位有效位
package main
import (
		"fmt"
		)

func main(){
	var sum1,sum2,sum3 int
	var average float32
	fmt.Println("请输入3个整数：")
	fmt.Scanf("%d, %d, %d", &sum1, &sum2, &sum3)
	fmt.Println("计算平均数...")
	average = float32(sum1+sum2+sum3) /3.0
	fmt.Printf("输出计算结果：%6.2f\n", average)

}
