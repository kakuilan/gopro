//continue的使用,把10~30之间能被3整除的数输出
package main
import "fmt"

func main(){
	var n int
	for n=10;n<30;n++{
		if n%3!=0{
			continue
		}
		fmt.Printf(" %d ",n)
	}
	fmt.Println()

}
