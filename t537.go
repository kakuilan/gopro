//使用1,2,3,4这4个数字,组成互不重复的三位数,并按照一行5个输出
package main
import (
		"fmt"
		"strconv"
		)

func main(){
	//整型变量num用来存储生成的三位数
	var num,col int
	//字符串numStr用来存数位
	var numStr string

	//三重循环嵌套实现
	for i:=1;i<5;i++{
		for j:=1;j<5;j++{
			for k:=1;k<5;k++{
				//确保个位、十位、百位互不相同
				if i!=k && i!=j && j!=k {
					numStr = strconv.Itoa(i) + strconv.Itoa(j) + strconv.Itoa(k)
					num,_ = strconv.Atoi(numStr)
					fmt.Printf("  %d",num)
					col++
					//控制每行输出5个数
					if col==5{
						fmt.Printf("\n")
						col = 0
					}


				}
			
			}
		
		}
	
	}
	fmt.Println()
}
