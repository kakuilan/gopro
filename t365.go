//; 需要注意, method value会复制 receiver
package main
import "fmt"

type User struct {
	id int
	name string
}

func (self User) Test(){
	fmt.Println(self)
}

func main(){
	u := User{1, "Tom"}
	mValue := u.Test //立即复制receiver,因为不是指针类型,不受后续修改影响

	u.id, u.name = 2, "Jack"
	u.Test()
	mValue()


}
