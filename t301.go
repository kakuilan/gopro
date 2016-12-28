// 建议改用引用类型,其底层数据不会被复制
package main
import "fmt"

func main(){
    s := []int {1,2,3,4,5} //切片
	for i,v := range s { //复制 struct slice {pointer, len, cap}
		if i==0 {
			s = s[:3] //对slice的修改,不会影响range
			s[2] = 100
		}
		println(i,v)
	}
	fmt.Println(s)

}
