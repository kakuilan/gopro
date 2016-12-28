//任意类型变参函数的声明和调用
package main
import "fmt"

func main(){
	f1(2,"Go",8,"Language",'a',false,'A',3.14)

}

func f1(args ... interface{}){
	var num = make([]int, 0, 6)
	var str = make([]string, 0,6)
	var ch = make([]int32, 0,6)
	var other = make([]interface{}, 0,6)

	for _,arg := range args{
		switch v:= arg.(type) {
			case int:
				num = append(num, v)
			case string:
				str = append(str, v)
			case int32:
				ch = append(ch, v)
			default:
				other = append(other, v)
		}
	}

	fmt.Println(num)
	fmt.Println(str)
	fmt.Println(ch)
	fmt.Println(other)
}
