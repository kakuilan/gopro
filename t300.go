// 注意,range会复制对象
package main
import "fmt"

func main(){
	a := [3]int {0,1,2}
	for i,v := range a { //index,value都是从复制品中取出
		if i == 0{
			a[1],a[2] = 888,999
			fmt.Println(a)
		}
		a[i] = v+100
	}
	fmt.Println(22, a)

}
