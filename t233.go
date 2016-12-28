// 创建一个固定大小保存整数，且后进先出LIFO的栈
// push函数将数据放入栈,pop函数从栈取数据,String函数将栈转化为字符串形式表达

package main
import (
    "fmt"
    "strconv"
)

//先定义一个新的类型来表达栈;需要一个数组来保存键和一个指向最后一个元素的索引.该栈只能保存10个元素
type stack struct {
    i int
    data [10]int
}

//使用指针,否则将只是s的副本
func (s *stack) push(k int){
    if s.i+1 > 9 {
	return
    }
    s.data[s.i] = k
    s.i++
}

func (s *stack) pop() int{
    s.i--
    return s.data[s.i]
}

func (s stack) String() string {
    var str string
    for i:=0;i<=s.i;i++{
	str = str + "[" +
		strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
    }
    return str
}

func main(){
    var s stack //让s是一个stack变量
    s.push(24)
    fmt.Println(s)
    s.push(21)
    fmt.Println(s)
    str := s.String()
    fmt.Println(str)

}
