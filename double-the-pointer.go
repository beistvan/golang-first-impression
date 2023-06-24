package main

import "fmt"

func doubleThePointer(a *int) {
	*a = *a * 2
}

func main() {
	b := 3
	doubleThePointer(&b)
	fmt.Println(b)
}
