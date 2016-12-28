//有一个3x4的整型矩阵,找出值最大的元素,并输出行号和列号
package main
import "fmt"

func main(){
	var i,j,row,col,max int
	var a = [3][4]int{{1,7,3,12},{4,9,5,19},{8,22,6,10}}
	max = a[0][0]

	for i=0;i<=2;i++{
		for j=0;j<=3;j++{
			if a[i][j] > max {
				max = a[i][j]
				row = i
				col = j
			}
		}
	}

	fmt.Printf("max=%d, row=%d, col=%d\n", max, row, col)
}
