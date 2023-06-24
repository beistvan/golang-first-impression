package main

import "fmt"

func main() {

	var num int

	fmt.Scanln(&num)

	prod := 1

	for i := 1; i <= num; i++ {
		prod *= i
	}

	fmt.Println(prod)
}
