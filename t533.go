//判断一个数是否为素数
//算法:要判断num是否素数,可让num除以2~k(k=int (num^(1/2)) )之间的数(num的平方根),如果num能被2~k之间的任何一个数整除,则num不是素数;否则,num是素数
package main
import (
		"fmt"
		"math"
		)

func main(){
	var num,i,k int
	fmt.Println("请输入一个整数：")
	fmt.Scanf("%d", &num)
	//对num开方并取整,求得k
	k = int(math.Sqrt(float64(num)))

	for i=2;i<=k;i++{
		if num % i == 0{
			break //num能被2~k中的一个数整除,跳出循环
		}
	}

	if i>k {
		fmt.Printf("%d是一个素数\n", num)
	}else{
		fmt.Printf("%d不是一个素数\n", num)
	}

}
