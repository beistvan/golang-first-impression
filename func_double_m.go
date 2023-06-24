package main

import "fmt"

func double_m(a, b int) int {
	s := 0
	for i := a; i <= b; i++ {
		s += i * i
	}
	return s
}

func main() {
	fmt.Println(double_m(2, 4))
}
