//for range语句对数据对象进行遍历
package main
import "fmt"

func main(){
	var str string = "Golang"
	//设k为键、v为值
	for k,v := range str {
		fmt.Printf("s[%d] = %c\n", k, v)
	}


}
