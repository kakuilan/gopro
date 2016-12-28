//使用函数嵌套调用,求正整数a,b的最小公倍数
//设a,b最大公约数为d,最小公倍数为c,则c=(a x b)/d
package main
import "fmt"

func main(){
	var a,b,c int
	fmt.Println("请输入正整数a和b:")
	fmt.Scanf("%d,%d", &a, &b)

	c = multiple(a,b)
	fmt.Printf("正整数%d和%d的最小公倍数为%d.\n", a,b,c)
}

//函数divisor()功能是求最大公约数
func divisor(a,b int) int{
	var r int
	for{
		r = a % b
		if r!= 0 {
			a = b
			b = r
		}else{
			break
		}
	}
	return b
}

//函数multiple的功能是求最小公倍数
func multiple(a,b int) int{
	var d int
	d = divisor(a,b)
	return (a*b /d)
}

