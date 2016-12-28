// 从文件读取(缓冲)
package main
import ("os";"bufio")

func main(){
    buf := make([]byte, 1024)
    f,_ := os.Open("/etc/passwd")
    defer f.Close()
    r := bufio.NewReader(f) //转换f为有缓冲的Reader
    w := bufio.NewWriter(os.Stdout)
    defer w.Flush()
    for {
	n,_ := r.Read(buf) //从Reader读取
	if n==0 {break}
	w.Write(buf[:n]) //而向Writer写入,然后向屏幕输出
    }


}
