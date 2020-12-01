//init函数
package mypkg

import "math"

var Pi float64

func init() {
	Pi = 4 * math.Atan(1)
}
