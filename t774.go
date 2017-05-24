//panic,recover
package main
import (
    "fmt"
    "math"
)

func ConvertInt64ToInt(x int64) int {
    if math.MinInt32 <=x && x <=math.MaxInt32 {
        return int(x)
    }
    //抛出异常
    panic(fmt.Sprintf("%d is out of the int32 range", x))
}

func IntFromInt64(x int64) (i int,err error) {
    //延迟处理
    defer func() {
        //恢复异常,转为一个错误返回
        if e := recover(); e!=nil {
            err = fmt.Errorf("%v", e)
        }
    }()
    i = ConvertInt64ToInt(x)
    return i,nil
}

func main() {
    var x int64
    x = 1<<31+1 
    d,err := IntFromInt64(x)
    fmt.Println(x,d,err)

}



