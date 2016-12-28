//1,2,3,4组成无重复数字的三位数
package main
import "fmt"

func main(){
	for i:=1;i<5;i++{
		for j:=1;j<5;j++{
			for k:=1;k<5;k++{
				//确保i,j,k互不相同
				if i!=k && i!=j && j!=k {
					fmt.Println(i,j,k)
				}
			}
		}
	}


}
