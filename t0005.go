package main

func main() {
	for i := 0; i < 5; i++ {
		println(i)
	}
	println("-------------------------")

	for i := 4; i >= 0; i-- {
		println(i)
	}
	println("-------------------------")

	x := 0
	for x < 5 {
		println(x)
		x++
	}
	println("-------------------------")

	y := 4
	for {
		println(x)
		y--
		if y < 0 {
			break
		}
	}
	println("-------------------------")

	s := []int{100, 101, 102}
	for i, n := range s {
		println(i, ":", n)
	}

}
