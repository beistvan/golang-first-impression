package main

import "fmt"

func main() {
	var gender int
	switch gender {
	case 1:
		fmt.Println("Man")
	case 2:
		fmt.Println("Woman")
	default:
		fmt.Println("In progress")
	}
}