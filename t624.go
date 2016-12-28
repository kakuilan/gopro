//defer语句保证文件能够正常关闭
package main
import (
		"fmt"
		"os"
		"io"
		)

func main(){
	copylen,err := copyFile("tmp624dst.txt", "tmp624src.txt")
	if err != nil{
		return
	}else{
		fmt.Println(copylen)
	}
}

//函数copyFile的功能是将源文件src的数据复制给目标文件dst
func copyFile(dstName,srcName string) (copylen int64,err error) {
	src,err := os.Open(srcName)
	if err != nil{
		return
	}
	//当return时就会调用src.Close()把源文件关闭
	defer src.Close()

	dst,err := os.Create(dstName)
	if err != nil{
		return
	}
	//当return时调用dst.Close()把目标文件关闭
	defer dst.Close()

	return io.Copy(dst,src)
}
