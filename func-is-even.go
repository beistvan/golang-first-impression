package main

import "fmt"

func isEven(a int) bool {
	if a%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isEven(4))
}
