// 执行命令
package main
import ("os/exec";"fmt")

func main(){
    cmd := exec.Command("/bin/ls", "-l")
    //err := cmd.Run() //执行命令
    buf,err := cmd.Output() //执行命令,并获取标准输出
    out := string(buf)
    fmt.Println(out, err)
}
