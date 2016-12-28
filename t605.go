//约瑟夫环问题
package main
import "fmt"

func main(){
	var psw [10]int
	var num,m,n,k int
	fmt.Println("请输入总人数n(<=10)和初始值m:")
	fmt.Scan(&n, &m)
	fmt.Println("请输入密码初始化数组...")
	for i := 0;i<n;i++{
		fmt.Scan(&psw[i])
	}
	fmt.Println("密码数组初始化完毕,开始报数...")

Label1 :
	for {
		for i:=0;i<n;i++{
			if psw[i] != 0{
				num++
				if num == m{
					m = psw[i]
					psw[i] = 0
					num = 0
					fmt.Printf(" %6d", i)
					k++
					if k==n {
						break Label1
					}
				}
			}
		}
	}

fmt.Println()

}
