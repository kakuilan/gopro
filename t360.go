//方法不过是一种特殊的函数,只需将其还原,就知道receiver T和*t的差别.
package main
import "fmt"

type Data struct {
	x int
}

//func ValueTest(self Data)
func (self Data) ValueTest(){
	fmt.Printf("Value: %p\n", &self)
}

func (self *Data) PointerTest(){
	fmt.Printf("Pointer:%p\n", self)
}

func main(){
	d := Data{}
	p := &d
	fmt.Printf("Data: %p\n", p)

	d.ValueTest() // ValueTest(d)
	d.PointerTest() // PointerTest(&d)

	p.ValueTest() // ValueTest(*p)
	p.PointerTest() // PointerTest(p)


}


