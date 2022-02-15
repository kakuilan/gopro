//二叉树-创建节点和设置数据
package main

import "fmt"

//使用空接口作为数据字段的类型
type Node struct {
	le   *Node
	data interface{}
	ri   *Node
}

func NewNode(left, right *Node) *Node {
	return &Node{left, nil, right}
}

func (n *Node) SetData(data interface{}) {
	n.data = data
}

func main() {
	root := NewNode(nil, nil)
	root.SetData("root node")

	//make child(leaf) nodes
	a := NewNode(nil, nil)
	a.SetData("left node")
	b := NewNode(nil, nil)
	b.SetData("right node")

	root.le = a
	root.ri = b
	fmt.Printf("%v\n", root)
}
