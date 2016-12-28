//链表综合操作
package main
import "link"

func main(){
	var head *link.Node
	stu1 := link.Node{link.Student{100,"李明"}, nil}
	stu2 := link.Node{link.Student{101,"张三"}, nil}
	stu3 := link.Node{link.Student{102,"赵武"}, nil}
	stu4 := link.Node{link.Student{103,"王丽"}, nil}

	//创建新链表
	head = head.Creat()
	
	//插入节点
	head = stu1.Insert(head)
	head = stu2.Insert(head)
	head = stu3.Insert(head)
	head = stu4.Insert(head)

	//输出链表
	head.PrintLink()
	//删除节点
	head = stu3.Delete(head)
	head.PrintLink()

}
