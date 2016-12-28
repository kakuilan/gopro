// GOMAXPROCS 设置可同时运行的线程数
package main
import (
    "fmt"
    "runtime"
)

func main(){
    n:= runtime.GOMAXPROCS(runtime.NumCPU())
    fmt.Println(n)
}
