package main

import "fmt"

func concat(a, b string) string {
	return a + b
}

func main() {
	fmt.Println(concat("Hello ", "World!"))
}
