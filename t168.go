// ioutil ReadAll
package main
import (
    "fmt"
    "io/ioutil"
    "os"
)

func main(){
    f,err := os.OpenFile("test.txt",os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
    if err != nil {
	fmt.Println(err.Error())
	return
    }

    defer f.Close()
    buf,rerr := ioutil.ReadAll(f)
    if rerr != nil {
	fmt.Println(rerr.Error())
	return
    }

    fmt.Println(string(buf))
}
