//考场座位有10排,每排5个,座位号0~49,按照S型排座位并打印出座位号
package main
import (
		"fmt"
		//"strconv"
		)

func main(){
	//定义并初始化座位为10行5列
	var row,col int=10,5
	//定义座位号
	var seatNo string
	for k:=0;k<row;k++{
		for i,j:=k*col,(k+1)*col-1;i<(k+1)*col && j>=k*col;i,j=i+1,j-1{
			if k%2 ==0{
				//seatNo+= " "+strconv.Itoa(i)
				seatNo += " " + fmt.Sprintf("%03d", i)
			}else{
				//seatNo+= " "+strconv.Itoa(j)
				seatNo += " " + fmt.Sprintf("%03d", j)
			}
		}
		fmt.Printf("%s\n", seatNo)
		seatNo = ""
	}


}
