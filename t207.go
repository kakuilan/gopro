// 模板函数 and or 
package main
import (
    "fmt"
    "html/template"
    "os"
)

func main(){
    strTpl := "and a b 结果为{{and .a .b}}\r\n or a b 结果为{{or .a .b}}\r\n"
    data := make(map[string] bool)
    data["a"] = true
    data["b"] = false
    t,err := template.New("test").Parse(strTpl)
    if err != nil {
	fmt.Println(err)
	return
    }
    err = t.Execute(os.Stdout, data)
    if err != nil {
	fmt.Println(err)
    }
}
