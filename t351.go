// struct 支持匿名结构,可用作结构成员或定义变量.
package main
import "fmt"

type File struct {
	name string
	size int
	attr struct {
		perm int
		owner int
	}
}

func main(){
	f := File{
		name : "test.txt",   
		size : 1025,
	}

	f.attr.owner = 1
	f.attr.perm = 0755

	var attr = struct {
		perm int
		owner int
	}{2, 0755}
	f.attr = attr

	fmt.Println(f)




}
