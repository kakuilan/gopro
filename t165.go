// path.Clean 返回跟path等价的短路径
package main
import (
    "fmt"
    "path"
    "path/filepath"
)

func main(){
    fmt.Println(path.Clean("/a/b/../c"))
    fmt.Println(path.Clean("/a/b/../././c"))

    //获取目录
    fmt.Println(path.Dir("/a/b/../c/d/e"))
    fmt.Println(path.Clean("/a/b/"))

    //获取扩展名
    fmt.Println(path.Ext("/a/b/../c/d./e"))
    fmt.Println(path.Ext("/a/b/test.txt"))

    //是否绝对路径
    fmt.Println(path.IsAbs("/a/b/c"))
    fmt.Println(path.IsAbs("/a/b/../c/e"))

    //路径连接
    fmt.Println(path.Join("a/b", "e"))

    //把路径分割为目录和文件两部分
    fmt.Println(path.Split("/a/b/test.txt"))
    fmt.Println(path.Split("/a/b/c/"))

    //把相对路径转换为绝对路径
    fmt.Println(filepath.Abs("."))

}
