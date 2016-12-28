// 延迟调用
package main
import (
	"fmt"
	"os"
)

func test() error {
	f,err := os.Create("test.txt")
	if err != nil {return err}

	defer f.Close() //注册调用,而不是注册函数.必须提供参数,哪怕为空

	f.WriteString("Hello, World! test 319.go\n")
	return nil
}

func main(){
	var r error
	r = test()
	fmt.Println(r)
}
