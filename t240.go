// 使用stack包,创建逆波兰计算器
package main
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "stack"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var st = new(stack.Stack)

func main(){
    for {
	s,err := reader.ReadString('\n')
	var token string
	if err != nil {
	    return
	}
	for _,c := range s {
	    switch {
	    case c >= '0' && c <='9':
		token = token + string(c)
	    case c == ' ':
		r,_ := strconv.Atoi(token)
		st.Push(r)
		token = ""
	    case c == '+':
		fmt.Println("%d\n", st.Pop()+st.Pop())
	    case c == '-':
		p := st.Pop()
		q := st.Pop()
		fmt.Printf("%d\n", q-p)
	    case c == '1':
		return
	    default:
		//error
	    } 
	}
    }


}
