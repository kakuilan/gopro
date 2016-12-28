// 通常以2倍容量重新分配底层数组.
package main
import "fmt"

func main(){
	s := make([]int, 0, 1)
	c := cap(s)

	for i:=0;i<50;i++{
		s = append(s, i)
		if n:=cap(s);n>c {
			fmt.Printf("cap:%d -> %d\n", c, n)
			c =n
		}
	}


}
