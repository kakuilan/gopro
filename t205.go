// template if else
package main
import (
    "fmt"
    "html/template"
    "os"
)

func main(){
    strTpl := "{{if .IsLogin}} 已登录 {{else}} 请登录 {{end}}\r\n{{if .IsVip}} 贵宾{{else}} 非贵宾{{end}}\r\n"
    data := make(map[string] bool)
    data["IsLogin"] = true
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
