//判断子串在主串中出现的位置
package main
import (
		"fmt"
		"strings"
		)

func main(){
	var s,sep,chars string = "Golang", "an", "lang"
	var r rune = 'a'
	
	fmt.Printf("%q第一次在%q中出现的索引是：%d\n", sep,s, strings.Index(s,sep))
	fmt.Printf("%q中的字符第一次在%q中出现的索引是：%d\n", chars,s, strings.IndexAny(s,chars))
	fmt.Printf("%q第一个符合f()的字符索引是：%d\n", s, strings.IndexFunc(s,f))
	fmt.Printf("%q第一次在%q中出现的索引是：%d\n", r, s, strings.IndexRune(s, r))
	fmt.Printf("%q最后一次在%q中出现的索引是：%d\n", sep,s, strings.LastIndex(s,sep))
	fmt.Printf("%q中的字符最后一次在%q中出现的索引是：%d\n", chars, s, strings.LastIndexAny(s, chars))
	fmt.Printf("%q最后一个符合f()的字符索引是：%d\n", s, strings.LastIndexFunc(s,f))

}

func f(a rune) bool {
	if a >'k'{
		return true
	}else{
		return false
	}
}
