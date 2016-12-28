//使用math包查看整型数据的最大、最小值
//可以检查数据溢出
package main
import(
		"fmt"
		"math"
		)

func main(){
	fmt.Printf("Min(int8)= %v, Max(int8)= %v\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("Min(int16)= %v, Max(int16)= %v\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("Min(int32)= %v, Max(int32)= %v\n", math.MinInt32, math.MaxInt32)


}
