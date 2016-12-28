// 从byte/int slice 到string 的转换
package main
import "fmt"

func main(){
	var str string
	var byteSlice = []byte{'H','e','l','l','o'}
	str = string(byteSlice)
	fmt.Println(byteSlice)
	fmt.Println(str)

}
