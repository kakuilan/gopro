package main
import (
    "os"
    "text/template"
)

var (
    zniota []rune
    fmap = template.FuncMap{
	"znmap" : znmap,
    }
)

func init(){
    zniota = []rune("壹贰叁肆伍陆柒捌玖拾")
}

func znmap(i int) rune {
    return zniota[i]
}

const tpl = `
{{range $i,$v :=.}}
{{$i|znmap|printf "%c"}}: {{$v.A姓名}}
{{end}}
`

type 剑客 struct {
    A姓名, B级别, C性别, D爱好 string
}

func main(){
    var 剑客录 = []剑客{
	{"骨惊飞", "北极", "男", "女"},
	{"骨灰飞", "", "有", "无"},
    }

    tmpl := template.New("")
    tmpl.Funcs(fmap)
    template.Must(tmpl.Parse(tpl))
    tmpl.Execute(os.Stdout, 剑客录)
}
