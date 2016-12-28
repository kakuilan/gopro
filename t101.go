package main
import (
    "os"
    "text/template"
)

const tpl = `
{{range .}}
{{if .B级别}}{{.A姓名}}大爷请上座。
{{else}}{{.A姓名}}小弟，欢迎再来。
{{end}}
{{end}}
`
type 剑客 struct {
    A姓名, B级别, C性别, D爱好 string
}

func main(){
    var 剑客录 = []剑客{
	{"古剑飞", "北极", "男", "女"},
	{"骨灰分", "", "有", "无"},
    }

    tmpl := template.New("")
    template.Must(tmpl.Parse(tpl))
    tmpl.Execute(os.Stdout, 剑客录)
}
