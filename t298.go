// for循环
package main

func length(s string) int{
    println("call length.")
    return len(s)
}

func main(){
    s := "abcd"
    for i,n:=0,length(s);i<n;i++{ //避免多次调用length函数
	println(i, s[i])
    }

}
