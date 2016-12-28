//通过匿名字段类型前缀访问同层级的同名字段
package main
import "fmt"

type d1 struct {
	x int
}

type d2 struct {
	x int
}

type data struct {
	d1
	d2
}

func main(){
	d := data{d1{10}, d2{20}}
	
	//这样访问将出错
	//err:ambiguous selector d.x 即d.x访问歧义
	//fmt.Println(d.x)

	//正确的访问
	fmt.Println(d.d1.x, d.d2.x)
}
