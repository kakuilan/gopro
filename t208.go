// 模板自定义函数
package main
import (
    "fmt"
    "html/template"
    "os"
)

func main(){
    strTpl := "{{SayHello}}\r\n"
    funcs := make(template.FuncMap)
    funcs["SayHello"] = SayHello
    t,err := template.New("test").Funcs(funcs).Parse(strTpl)
    if err != nil {
	fmt.Println(err)
	return
    }

    err = t.Execute(os.Stdout, nil)
    if err != nil {
	fmt.Println(err)
    }
}

func SayHello()string {
    return "你好，自定义模板函数"
}
