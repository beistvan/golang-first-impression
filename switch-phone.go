package main

import "fmt"

func main() {

	var money int
	fmt.Scanln(&money)
	switch {
	case money > 1000:
		fmt.Println("Apple")
	case money >= 500 && money <= 1000:
		fmt.Println("Samsung")
	case money < 500:
		fmt.Println("Nokia with flashlight")
	}
}
