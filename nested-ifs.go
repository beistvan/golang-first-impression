package main

import "fmt"

func main() {

	var result = 1

	if result == 3 {
		fmt.Println("Win")
	} else if result == 1 {
		fmt.Println("Deuce")
	} else {
		fmt.Println("Defeat")
	}
}