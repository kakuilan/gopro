// 将方法"还原"成函数,就容易理解下面的代码了
package main
import "fmt"

type Data struct {}

func (Data) TestValue(){}
func (*Data) TestPointer(){}

func main(){
	var p *Data = nil
	p.TestPointer()

	(*Data)(nil).TestPointer() // method value
	(*Data).TestPointer(nil) //method expression

	fmt.Println(p)


}
