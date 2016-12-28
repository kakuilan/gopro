// stuct结构体,值类型,赋值和传参会赋值全部内容.可用"_"定义补位字段,支持指向自身类型的指针成员.
package main
import "fmt"

type Node struct {
	_ int
	id int
	data *byte
	next *Node
}

func main(){
	n1 := Node{
		id : 1,
		data : nil,
	}
	n2 := Node{
		id : 2,
		data : nil, 
		next : &n1,
	}

	fmt.Println(n1, n2)

}


