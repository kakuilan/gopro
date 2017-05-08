//反序列化一个其内部结构未知的原始JSON对象,类型断言
package main
import (
    "fmt"
    "encoding/json"
    "bytes"
)

func main() {
    MA := []byte(`{"name":"MassSCHutts","area":27336,"water":25.7,"senatiors":["John kerry","Scoot bron"]}`)
    var object interface{}
    if err := json.Unmarshal(MA, &object);err !=nil {
        fmt.Println(err)
    }else{
        jsonObject := object.(map[string]interface{})
        fmt.Println(jsonObjectAsString(jsonObject))
    }
}

//使用一个未检查的类型断言预语句
func jsonObjectAsString(jsonObject map[string]interface{}) string {
    var buffer bytes.Buffer
    buffer.WriteString("{")
    comma := ""
    for key,value := range jsonObject {
        buffer.WriteString(comma)
        //下面的value 是影子变量
        switch value:= value.(type) {
        case nil :
            fmt.Fprintf(&buffer, "%q:null", key)
        case bool :
            fmt.Fprintf(&buffer, "%q:%t", key, value)
        case float64:
            fmt.Fprintf(&buffer, "%q:%f", key, value)
        case string:
            fmt.Fprintf(&buffer, "%q:%q", key, value)
        case []interface {} :
            fmt.Fprintf(&buffer, "%q: [", key)
            innerComma := ""
            for _,s := range value {
                if s,ok := s.(string);ok {
                    fmt.Fprintf(&buffer, "%s%q", innerComma, s)
                    innerComma = ", "
                }
            }
            buffer.WriteString("]")
        }
        comma = ", "
    }
    buffer.WriteString("}")
    return buffer.String()
}
