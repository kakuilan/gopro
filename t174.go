// xml编码
package main
import (
    "encoding/xml"
    "fmt"
    "os"
)

type Student struct {
    Name string
    Age int
}

func main(){
    f,err := os.Create("174.dat")
    if err != nil {
	fmt.Println(err.Error())
	return
    }

    defer f.Close()
    s := &Student{"李四", 34}
    //创建encode对象
    encoder := xml.NewEncoder(f)
    //将s序列化到文件中
    encoder.Encode(s)
    //重置文件指针到开始位置
    f.Seek(0, os.SEEK_SET)
    decoder := xml.NewDecoder(f)
    var s1 Student
    //从文件中反序列化成对象
    decoder.Decode(&s1)
    fmt.Println(s1)





}

