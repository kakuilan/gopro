//设定谜底数字为36,设计一个猜数字游戏,提示玩家数字过大还是过小,最后统计玩家共猜了多少次才猜中
package main
import (
		"bufio"
		"fmt"
		"os"
		"strconv"
		)

const GOAL int = 36 //设置谜底

func main(){
	var data int //存储输入数字
	var count int //统计用了多少次才猜中
	var c byte //存放每个按键ASCII码
	var str string //存放输入字符串流

	fmt.Println("请输入一个正整数：")
LABEL1:
	r := bufio.NewReader(os.Stdin) //初始化标准输入设备
	w := bufio.NewWriter(os.Stdout) //初始化标准输出设备

	//键盘输入处理流程
	for i:=0;;i++{
		c,_ = r.ReadByte()
		if c==10||c==13{
			//如果是回车或换行,则退出
			break
		}else{
			w.Flush()
			str += string(c)
		}
	}

	//猜数字流程
	for{
		data,_ = strconv.Atoi(str)
		if data > GOAL{
			fmt.Printf("数字%d有点大,请再试一试...\n", data)
			count++
			str = ""
			goto LABEL1
		}else if data < GOAL {
			fmt.Printf("数字%d有点小,请再试一试...\n", data)
			count++
			str = ""
			goto LABEL1
		}else{
			fmt.Println("恭喜你，猜中了！", data)
			break
		}
	
	}

	fmt.Printf("你一共猜了%d次才猜对.\n", count)
}


