// package 例子
package stack

type Stack struct {
    i int
    data [10]int
}

func (s *Stack) Push(k int) {
    if s.i+1 > 9 {
	return
    }

    s.data[s.i] = k
    s.i++
}

func (s *Stack) Pop() (ret int) {
    s.i--
    if s.i <0 {
	s.i = 0
	return
    }
    ret = s.data[s.i]
    return
}
