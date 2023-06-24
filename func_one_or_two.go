package main

import "fmt"

func one_or_two(a, b int, s string) (int, string) {
	switch {
	case s == "one":
		return a, s
	case s == "two":
		return b, s
	default:
		return 0, s
	}
}

func main() {
	fmt.Println(one_or_two(2, 4, "one"))
	fmt.Println(one_or_two(2, 4, "two"))
	fmt.Println(one_or_two(2, 4, "ok"))
}
