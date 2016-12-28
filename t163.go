// path
package main
import (
    "fmt"
    "path"
    "strings"
)

func main(){
    fmt.Println(path.Base("/usr/bin"))
    fmt.Println(path.Base(""))
    fmt.Println(path.Base("C:\\Windows")) //无法识别windows下的路径分隔符
    fmt.Println(path.Base(strings.Replace("C:\\Windows","\\","/",-1))) //把\转换为/

}
