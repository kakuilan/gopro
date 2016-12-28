//使用for i:=0;;i++{}无线循环实现简单键盘文本输入
package main
import (
		"bufio"
		"os"
		"fmt"
		)

func main(){
	var c byte //存放每个按键ASCII码
	var str string //存放输入字符串流
	//初始化标准输入设备
	r := bufio.NewReader(os.Stdin)
	//初始化标准输出设备
	w := bufio.NewWriter(os.Stdout)

	for i:=0;;i++{
		//从标准输入设备接收字符
		c,_ = r.ReadByte()
		if c==10 {
			//如果是回车键,停止接收退出循环
			break
		}else{
			//将接收到的字符写入标准输出设备
			w.WriteByte(c)
			w.Flush()
			//生成字符串
			str += string(c)
		}
	}

	fmt.Printf("\n")
	fmt.Println(str)
	//将输出两行
	//第一行是从标准输入设备(键盘)上接收的字符,从标准输出设备(显示终端)输出
	//第二行是存放在字符串str中的数据,二者是一致的

}
