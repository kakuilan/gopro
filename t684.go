//二叉树的建立
package main
import (
		"fmt"
		. "btnode" //以此调用省略前缀的包名
		)

func main(){
	//创建根节点
	root := NewNode(nil,nil)
	fmt.Println(root)
	var it Initer
	it = root
	it.SetData("root node")

	//创建左子树
	a := NewNode(nil,nil)
	a.SetData("left node")
	al := NewNode(nil,nil) //左叶子节点
	al.SetData(100)
	ar := NewNode(nil,nil) //右叶子节点
	ar.SetData(3.14)
	a.Left = al
	a.Right = ar

	//创建右子树
	b := NewNode(nil,nil)
	b.SetData("right node")
	root.Left = a
	root.Right = b
	root.PrintBT()
	fmt.Println()



}

