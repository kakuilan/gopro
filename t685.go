//二叉树的基本操作
package main
import (
		"fmt"
		. "btnode"
		)

func main(){
	//创建二叉树
	root := NewNode(nil,nil)
	root.SetData("root node")
	a := NewNode(nil,nil)
	a.SetData("left node")
	al := NewNode(nil,nil)
	al.SetData(100)
	ar := NewNode(nil,nil)
	ar.SetData(3.14)
	a.Left = al
	a.Right = ar
	b := NewNode(nil,nil)
	b.SetData("right node")
	root.Left = a
	root.Right = b

	//使用Operater接口实现对二叉树的基本操作
	var it Operater
	it = root
	it.PrintBT()
	fmt.Println()
	fmt.Println("The depths of the Btree is:", it.Depth())
	fmt.Println("The leaf counts of the Btree is:", it.LeafCount())

}
