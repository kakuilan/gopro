//模板包2
package main
import (
    "os"
    "text/template"
)

const tpl = `
知止{{range .}}而后能{{.}},{{.}}{{end}}而后能得.
`

func main(){
    var 大学 = []string{"定", "静", "安", "虑"}
    tmpl := template.New("")
    tmpl.Parse(tpl)
    tmpl.Execute(os.Stdout, 大学)
}

