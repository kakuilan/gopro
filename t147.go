// CPU核数
package main
import (
    "fmt"
    "runtime"
)

func main(){
    fmt.Println(runtime.NumCPU())
    fmt.Println(runtime.NumGoroutine())
}
