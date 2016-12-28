// 文件读写
package main
import (
    "fmt"
    "io"
    "os"
)

func main(){
    //打开test.txt文件,如果文件不存在则创建,若存在则追加内容到文件尾
    f,err := os.OpenFile("test.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
    if err != nil {
	fmt.Println(err.Error())
	return
    }

    r,werr :=f.WriteString("\r\n文件测试\r\n试试看\r\n")
    fmt.Println(r,werr)
    //f.Close()
    //return
    buf := make([]byte, 1024)
    var str string

    //重置文件指针,否则读不到内容
    f.Seek(0, os.SEEK_SET)
    for {
	n,ferr := f.Read(buf)
	if ferr != nil && ferr != io.EOF {
	    fmt.Println(ferr.Error())
	    break
	}
	if n==0 {
	    break
	}
	fmt.Println(str)
	str += string(buf[0:n])
    }

    fmt.Println(str)
    f.Close()

}
