//可直接修改struct array/slice成员
package main
import "fmt"

func main(){
	d := [5] struct {
		x int   
	}{}

	s := d[:]
	d[1].x = 10
	fmt.Println(d)
	s[2].x = 20
	fmt.Println(s)
	fmt.Printf("%p, %p \n", &d, &d[0])


}
