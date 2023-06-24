package main

import "fmt"

type Car struct {
	color, brand string
	year         int
}

func main() {
	x := Car{"Volvo", "blue", 2023}
	fmt.Println(x.color)
	fmt.Println(x.brand)
	fmt.Println(x.year)
}
