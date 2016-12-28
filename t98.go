//模板包
package main
import (
    "os"
    "text/template"
)

const tpl = `
How many roads must {{.}} walk down
Before they call him {{.}}
`

func main(){
    tmpl := new(template.Template)
    tmpl.Parse(tpl)
    tmpl.Execute(os.Stdout, "a man")
}
