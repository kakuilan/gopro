//利用Comma-ok断言进行接口类型推断
package main
import "fmt"

type People struct {
	Name string
	Age int
}

//定义空接口用于存储任意数据类型
type Tester interface{}

func main(){
	people := People{"张三", 20}
	it := make([]Tester, 4)
	it[0],it[1],it[2],it[3] = 1,"Hello",people,true

	for index,element := range it {
		if value,ok := element.(int);ok{
			fmt.Printf("it[%d] is an int. value=%d\n", index, value)
		}else if value,ok := element.(string);ok{
			fmt.Printf("it[%d] is a string. value=%s\n", index, value)
		}else if value,ok := element.(People);ok{
			fmt.Printf("it[%d] is a People. value=%v\n", index, value)
		}else{
			fmt.Printf("it[%d] is an unknown type.\n", index)
		}
	}


}
