// 整型数据的溢出
//int8最大为127,超出最大值就会发生溢出
package main
import "fmt"

func main(){
	var a int8 = 127
	fmt.Println(a, a+1) //-128

}
