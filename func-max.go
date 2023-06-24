package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func main() {
	fmt.Println(max(2, 3))
}
