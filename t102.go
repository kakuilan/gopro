package main
import (
    "os"
    "text/template"
)

const tpl = `
{{range $i,$v := .}}{{$i}}:{{$v.A姓名}}
{{end}}
`

type 剑客 struct {
    A姓名, B级别, C性别, D爱好 string
}

func main(){
    var 剑客录 =[]剑客{
	{"骨惊飞", "北极", "男", "女"},
	{"骨灰飞", "", "有", "无"},
    }

    tmpl := template.New("")
    template.Must(tmpl.Parse(tpl))
    tmpl.Execute(os.Stdout, 剑客录)
}
