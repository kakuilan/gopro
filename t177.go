// xml UnMarshal xml字段小写处理
package main
import (
    "encoding/xml"
    "fmt"
)

type Student struct {
    XMLName xml.Name `xml:"student"`
    Name string `xml:"name"`
    Age int `xml:"age"`
}

func main(){
    str := `
	<?xml version="1.0" encoding="utf-8" ?>
	<student>
	<name>展昭</name>
	<age>29</age>
	</student>`
    var s Student
    xml.Unmarshal([]byte(str), &s)
    fmt.Println(s)

}
