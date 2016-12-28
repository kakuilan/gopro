//正则,小行星命名
package main
import (
    "fmt"
    "io/ioutil"
    "regexp"
)

var pat = `\d{4}\s*[A-HJ-Y][A-HJ-Z]\d*`

func main(){
    text,_ := ioutil.ReadFile("regexp.t")
    reg,_ := regexp.Compile(pat)
    for i,v := range reg.FindAll(text, -1) {
	fmt.Printf("%d : %s\n", i, v)
    }
}

