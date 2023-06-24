package main

import "fmt"

type Car struct {
	color, brand string
	year         int
}

func main() {
	volvo := Car{"Volvo", "blue", 2023}
	fmt.Println(volvo.color)
	fmt.Println(volvo.brand)
	fmt.Println(volvo.year)

	porshe := Car{"Porshe", "red", 1991}
	fmt.Println(porshe.color)
	fmt.Println(porshe.brand)
	fmt.Println(porshe.year)
}
