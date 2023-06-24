package main

import "fmt"

func mars_age(a int) int {
	return a * 365 / 687
}

func main() {
	fmt.Println(mars_age(30))
}
