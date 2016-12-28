// ioutil ReadDir
package main
import (
    "fmt"
    "io/ioutil"
)

func main(){
    arrFile,err := ioutil.ReadDir(".")
    if err != nil {
	fmt.Println(err.Error())
    }

    for k,v := range arrFile {
	fmt.Println(k, "\t", v.Name(), "\t", v.IsDir())
    }
}
