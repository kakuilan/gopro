// 滥用defer可能会导致性能问题,尤其是在一个大循环里
package t321
import (
	"testing"
	"sync"
)

var lock sync.Mutex

func test(){
	lock.Lock()
	lock.Unlock()
}

func testdefer(){
	lock.Lock()
	defer lock.Unlock()
}

func BenchmarkTest(b *testing.B){
	for i:=0;i<b.N;i++{
		test()
	}
}

func BenchmarkTestDefer(b *testing.B){
	for i:=0;i<b.N;i++{
		testdefer()
	}
}

