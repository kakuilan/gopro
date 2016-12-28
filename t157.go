// 读写互斥
package main
import (
    "errors"
    "fmt"
    "sync"
)

type MyMap struct {
    mp map[string]int
    mutex *sync.RWMutex
}

func (this *MyMap) Set(key string,v int) {
    this.mutex.Lock()
    defer this.mutex.Unlock()
    this.mp[key] = v
}

func (this *MyMap) Get(key string) (int,error) {
    this.mutex.RLock()
    i,ok := this.mp[key]
    this.mutex.RUnlock()
    if !ok {
	return i,errors.New("不存在")
    }
    return i,nil
}

func (this *MyMap) Display(){
    this.mutex.RLock()
    defer this.mutex.RUnlock()
    for k,v := range this.mp {
	fmt.Println(k,"=",v)
    }
}

func SetValue(m *MyMap) {
    var a rune
    a = 'a'
    for i:=0;i<10;i++{
        m.Set(string(a+rune(i)), i)
    }
}

func main(){
    m := &MyMap{mp:make(map[string]int), mutex:new(sync.RWMutex)}
    go SetValue(m) //启动一个线程向map写入值
    go m.Display() //启动一个线程读取map的值
    var str string //这里主要是等待线程结束
    fmt.Scan(&str)
    fmt.Println(m)
}
