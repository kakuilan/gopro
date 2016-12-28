// json Unmarshal
package main
import (
    "encoding/json"
    "fmt"
)

func main(){
    str := `{"userName":"呵呵","Age":20}`
    var m map[string] interface{}
    json.Unmarshal([]byte(str), &m)

    for k,v := range m {
	switch v.(type) {
	case float64:
	    fmt.Println(k, " 是int类型,值为:", v)
	case string:
	    fmt.Println(k, " 是string类型,值为:", v)
	default:
	    fmt.Println(k, "无法误用别的类型")
	}
    }

}
