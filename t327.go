//标准库 errors.New 和 fmt.Errorf函数用于创建实现error接口的错误对象.通过判断错误对象实例来确定具体错误类型
package main
import "errors"

var ErrDivByZero = errors.New("division by zero")

func div(x,y int) (int, error) {
	if y==0 {return 0, ErrDivByZero}
	return x/y, nil
}

func main(){
	switch z,err := div(10, 0);err{
		case nil :
			println(z)
		case ErrDivByZero:
			panic(err)
	}

}
