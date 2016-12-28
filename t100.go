package main
import (
    "os"
    "text/template"
)

const tpl = `
<tr>
<td>{{.A姓名}}</td>
<td>{{.B级别}}</td>
<td>{{.C性别}}</td>
<td>{{.D爱好}}</td>
</tr>
`

//注意下面 A中文 是没有空格的,否则出错 
type 剑客 struct {
    A姓名, B级别, C性别, D爱好 string
}

func main(){
    var 剑客录 = 剑客{"骨惊飞", "北极", "男", "女"}
    tmpl := template.New("")
    tmpl.Parse(tpl)
    tmpl.Execute(os.Stdout, 剑客录)
}
