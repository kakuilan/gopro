// 使用panic
package main
import "os"
import "fmt"

var user = os.Getenv("USER")

func init(){
	if user == ""{
		panic("no value for $USER")
	}
	fmt.Println(user)
}

//下面的函数检查作为其参数的函数在执行时是否会产生panic
func throwsPanic(f func()) (b bool) {
	defer func(){
		if x := recover(); x!=nil {
			b = true
		}
	}()

	f() //执行函数f,如果f中出现了panic,那么就可以恢复回来
	return
}

func kh(){
	panic("硬是要制造恐慌！")
}

func main(){
	println()
	kh()
}
