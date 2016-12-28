//获取文件路径

package main
import (
    "os"
    "log"
    "os/exec"
    "path"
    "path/filepath"
)

func main(){
    cf,_ := exec.LookPath(os.Args[0])
    cp,_ := filepath.Abs(cf)
    log.Println("当前文件:",cf)
    log.Println("当前路径:",cp)

    file,_ := os.Getwd()
    log.Println("当前路径:", file)

    file,_ = exec.LookPath(os.Args[0])
    log.Println("执行路径:", file)

    dir,_ := path.Split(file)
    log.Println("执行目录相对路径:", dir)

    os.Chdir(dir)
    wd,_ := os.Getwd()
    log.Println("执行目录绝对路径:",wd)



}
