// 测试包
package even
import "testing"

func TestEven(t *testing.T) {
    if Even(2) {
	t.Log("2 should be odd!")
	t.Fail()
    }
}
