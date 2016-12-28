// gob化
package main
import (
    "encoding/gob"
    "fmt"
    "os"
)

type Student struct {
    Name string
    Age int
}

func main(){
    s := &Student{"张三", 19}
    f,err := os.Create("173.dat")
    if err != nil {
	fmt.Println(err.Error())
	return
    }

    defer f.Close()

    //创建gob Encoder对象
    encode := gob.NewEncoder(f)
    //将s序列化到f文件中
    encode.Encode(s)
    //重置文件指针到开始位置
    f.Seek(0, os.SEEK_SET)
    decoder := gob.NewDecoder(f)
    var s1 Student
    //反序列化对象
    decoder.Decode(&s1)
    fmt.Println(s1)

}

