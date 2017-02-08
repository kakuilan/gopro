//time.Sleep函数,表示休眠多少时间,休眠时处于堵塞状态,后续程序无法执行
//time.AfteriFunc自定义定时器
package main
import (
    "fmt"
    "time"
)

func main(){
    var t *time.Timer

    f := func(){
        fmt.Printf("Expiration time: %v.\n", time.Now())
        fmt.Printf("C`s len: %d\n", len(t.C))
    }

    t = time.AfterFunc(1 * time.Second, f)
    
    //让当前goroutine睡眠2s,确保大于内容的完整
    //这样做原因是time.AfterFunc的调用不会被阻塞.它会以异步的方式在到期事件来临时执行我们的自定义函数f
    time.Sleep(2 * time.Second)

}
