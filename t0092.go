//获取内存状态
package main

import (
	"fmt"
	"runtime"
)

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("mem: %d Kb\n", m.Alloc/1024)

}
