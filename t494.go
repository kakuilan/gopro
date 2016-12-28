//整型数据之间的转换
package main
import "fmt"

func main(){
	var a int =128
	var b int8 =127
	var c int16 =32767
	var d int32 =32768
	var e int64 =32768

	fmt.Println(a, int8(a), int16(a), int32(a), int64(a))
	fmt.Println(b, int8(b), int16(b), int32(b), int64(b))
	fmt.Println(a, int8(c), int16(c), int32(c), int64(c))
	fmt.Println(a, int8(d), int16(d), int32(d), int64(d))
	fmt.Println(a, int8(e), int16(e), int32(e), int64(e))

}

