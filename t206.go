// template range
package main
import (
    "fmt"
    "html/template"
    "os"
)

func main(){
    strTpl := "{{range .test}}{{.}}\r\n{{end}} {{range .test1}}{{.}}\r\n{{else}}test1不存在.{{end}}\r\n"
    data := make(map[string] interface{})
    arr := []int{1,2,3,4}
    data["test"] = arr
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
