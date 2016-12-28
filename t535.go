//goto使用.计算1~100的累加和
package main
import "fmt"

func main(){
	var i,sum = 1,0
HERE: //标签
	sum = sum+i
	i++
	if i<=100{
		goto HERE
	}
	fmt.Println(sum)

}
