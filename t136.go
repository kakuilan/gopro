//defer例子
package main
import (
    "fmt"
    "os"
)

func ReadFile(strFileName string) (string, error) {
    f,err := os.Open(strFileName)
    if err != nil {
	fmt.Println()
	return "",err
    }

    defer f.Close() //在函数返回前关闭文件
    buf := make([]byte, 1024)
    var strContent string = ""
    for {
	n,_ := f.Read(buf)
	if n==0 {
	    break
	}
	strContent += string(buf[0:n])
    }

    return strContent,nil
}

func main(){
    str,err := ReadFile("t136.go")
    if err != nil {
	fmt.Println(err.Error())
	return
    }
    fmt.Println(str)

}
