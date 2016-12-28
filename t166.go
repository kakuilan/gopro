// 遍历目录
package main
import (
    "fmt"
    "os"
    "path/filepath"
)

func DispFile(path string, info os.FileInfo, err error) error {
    fmt.Println(path,"------",info.Name(),"------",info.IsDir())
    return nil
}

func main(){
    filepath.Walk(".", DispFile)
}
