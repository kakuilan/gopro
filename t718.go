// 浮点数精确度
package main

import (
	"fmt"
	"math"
)

func main() {
	x, y := 0.0, 0.0
	for i := 0; i < 10; i++ {
		x += 0.1
		if i%2 == 0 {
			y += 0.2
		} else {
			// %-5以一个向左对齐的5个字符宽的区域打印一个布尔值
			fmt.Printf("%-5t %-5t %-5t %-5t", x == y, EqualFloat(x, y, -1), EqualFloat(x, y, 0.00000000001), EqualFloatPrec(x, y, 6))
			fmt.Println(x, y)
		}
	}

	fmt.Println()
	var f float64
	var iv int
	f = 2.61001
	iv = IntFromFloat64(f)
	fmt.Println(f, iv)

}

func EqualFloat(x, y, limit float64) bool {
	if limit <= 0.0 {
		limit = math.SmallestNonzeroFloat64
	}
	return math.Abs(x-y) <= (limit * math.Min(math.Abs(x), math.Abs(y)))
}

func EqualFloatPrec(x, y float64, decimals int) bool {
	a := fmt.Sprintf("%.*f", decimals, x)
	b := fmt.Sprintf("%.*f", decimals, y)
	return len(a) == len(b) && a == b
}

func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole, fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%g is out of the int32 range", x))
}
