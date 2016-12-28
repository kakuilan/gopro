//strconv.ParseInt函数
package main
import (
	"strconv"
)

func main(){
	i,err := strconv.ParseInt("123", 10, 32)
	if err != nil {
		panic(err)
	}
	println(i)
}
