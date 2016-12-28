// 创建目录
package main
import ("os";"fmt")

func main(){
    dirname := "testdir"
    if f,e := os.Stat(dirname);e !=nil {
	os.Mkdir(dirname, 0755)
    }else{
	fmt.Println(f,e, dirname)
    }


}
