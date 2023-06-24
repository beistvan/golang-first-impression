package main

import "fmt"

func my_print(nums ...int) {
	for _, v := range nums {
		fmt.Println(v)
	}
}

func main() {
	numbers := []int{0, 2, 4, 6, 8}
	fmt.Println(numbers[1])
	my_print(numbers...)
	my_print(7, 3, 8, 2)
	my_print(2, 6, 5)
}
