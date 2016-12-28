// tcp test 服务端
package main
import (
    "bytes"
    "fmt"
    "io/ioutil"
    "net"
    "os"
)

const (
    LS = "LS"
    CD = "CD"
    PWD = "PWD"
)

func main(){
    //在7070端口监听
    tcpAddr,err := net.ResolveTCPAddr("tcp", ":7070")
    checkError(err)
    listener,err1 := net.ListenTCP("tcp", tcpAddr)
    checkError(err1)
    fmt.Println("tcp服务端已启动...")
    for {
	//等待客户端的连接
	conn,err2 := listener.Accept()
	if err2 != nil {
	    //通常服务端为一个服务,不会因为错误而退出.出错后,继续等待下一个连接请求
	    fmt.Println(err2)
	    continue
	}
	fmt.Println("收到客户端的请求")
	go ServeClient(conn)
    }
}

func ServeClient(conn net.Conn) {
    defer conn.Close()
    str := ReadData(conn)
    if str == "" {
	SendData(conn,"接受数据时出错")
	return
    }
    fmt.Println("收到命令:", str)
    switch str {
    case LS:
	ListDir(conn)
    case PWD:
	Pwd(conn)
    default:
	if str[0:2] == CD {
	    Chdir(conn, str[3:])
	}else{
	    SendData(conn, "命令错误")
	}
    }
}

func Chdir(conn net.Conn, s string) {
    err := os.Chdir(s)
    if err != nil {
	SendData(conn, err.Error())
    }else{
	SendData(conn, "OK")
    }
}

func ListDir(conn net.Conn){
    files,err := ioutil.ReadDir(".")
    if err != nil {
	SendData(conn, err.Error())
	return
    }
    var str string
    for i,j:=0,len(files);i<j;i++{
	f := files[i]
	str += f.Name() + "\t"
	if f.IsDir() {
	    str += "dir\r\n"
	}else{
	    str += "file\r\n"
	}
    }
    SendData(conn, str)
}

//读取数据
func ReadData(conn net.Conn) string {
    var data bytes.Buffer
    var buf [512]byte
    for {
	n,err := conn.Read(buf[0:])
	if err != nil {
	    fmt.Println(err)
	    return ""
	}
	//我们的数据以0作为结束的标记
	if buf[n-1]==0 {
	    //n-1去掉结束标记0
	    data.Write(buf[0:n-1])
	    break
	}else{
	    data.Write(buf[0:n])
	}
    }

    return string(data.Bytes())
}

//发送数据
func SendData(conn net.Conn, data string) {
    buf := []byte(data)
    //向byte字节里添加结束标记
    buf = append(buf, 0)
    _, err := conn.Write(buf)
    if err != nil {
	fmt.Println(err)
    }
} 

func Pwd(conn net.Conn) {
    s,err := os.Getwd()
    if err != nil {
	SendData(conn, err.Error())
    }else{
	SendData(conn, s)
    }
}

func checkError(err error) {
    if err != nil {
	fmt.Println(err)
	os.Exit(0)
    }
}
