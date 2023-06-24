package main

import "fmt"

func calc(a int) (int, int) {
	return a * 2, a * a
}

func main() {
	a, b := calc(3)
	fmt.Println(a + b)
	fmt.Println(calc(3))
}
