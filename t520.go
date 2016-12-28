//使用if嵌套语句判断闰年
package main
import "fmt"

func main(){
	var year int
	fmt.Println("请输入年份")
	fmt.Scanln(&year)

	if year % 4 == 0 {
		if year % 100 ==0{
			if year % 400==0 {
				//能被4整除,且能被100整除,且能被400整除的是闰年
				fmt.Printf("%d是闰年.\n", year)
			}else{
				//能被4、100整除,但不能被400整除的不是闰年
				fmt.Printf("%d不是闰年.\n", year)
			}
		}else{
			//能被4整除,但不能被100整除的是闰年
			fmt.Printf("%d是闰年.\n", year)
		}
	}else{
		//不能被4整除的不是闰年
		fmt.Printf("%d 不是闰年.\n", year)
	}


}
