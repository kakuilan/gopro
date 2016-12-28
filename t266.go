// 列出所有正在运行的进程,并打印每个进程执行的子进程个数
// 题解:
// 1.运行ps获取输出
// 2.解析输出并保存每个PPID的子PID
// 3.排序PPID列表
// 4.打印排序后的列表到屏幕
package main
import (
    "fmt"
    "os/exec"
    "sort"
    "strconv"
    "strings"
)

func main(){
    ps := exec.Command("ps", "-e", "-opid,ppid,comm")
    output,_ := ps.Output()
    child := make(map[int][]int)
    for i,s := range strings.Split(string(output), "\n"){
	if i==0 {continue} //第一行,如  PID  PPID COMMAND,不要
	if len(s)==0 {continue} //最后一行,PS,也不要
	f := strings.Fields(s)
	fpp,_ := strconv.Atoi(f[1]) //父pid
	fp,_ := strconv.Atoi(f[0]) //子pid
	child[fpp] = append(child[fpp], fp)
    }

    schild := make([]int, len(child))
    i := 0
    for k,_ := range child {schild[i]=k;i++}
    sort.Ints(schild)
    for _,ppid := range schild {
	fmt.Printf("Pid %d has %d child", ppid, len(child[ppid]))
	if len(child[ppid]) ==1 {
	    fmt.Printf(": %v\n", child[ppid])
	    continue
	}
	fmt.Printf("ren: %v\n", child[ppid])
    }

}

