// map. 预先给make函数一个合理元素数量参数,有助于提升性能.因为事先申请一大块内存,可避免后续操作时频繁扩张.
package main
import "fmt"

func main(){
	m0 := make(map[string]int, 1000)
	println(len(m0))
	
	//常见操作
	m := map[string]int{
		"a" : 1,
	}

	//判断key是否存在
	if v,ok := m["a"];ok {
		println(v,ok)
	}

	println(m["c"]) //对于不存在的key,直接返回\0,不会出错

	//新增或修改
	m["b"] = 2
	m["d"] = 3
	fmt.Println(m)

	//删除,如果key不存在,不会出错
	delete(m, "h")

	//获取键值对数量.cap无效
	println(len(m))

	//迭代,可仅返回key.随机顺序返回,每次都不相同
	for k,v := range m {
		println(k, v)
	}



}
