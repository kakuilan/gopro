//iota
package main
import "fmt"

const(
		a = 'A'
		b
		c = iota
		d
		)

const(
		e='E'
		f=iota
		)

func main(){
fmt.Println(a)
fmt.Println(b)
fmt.Println(c)
fmt.Println(d)
fmt.Println(e)
fmt.Println(f) 
}
