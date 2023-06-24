package main

import "fmt"

func main() {
	// nums := [5]int{0, 2, 4, 6, 8} // sum = 20
	nums := make([]int, 2)
	nums[0] = 1
	nums[1] = 5

	sum := 0

	for _, v := range nums {
		sum += v
	}

	fmt.Println(sum)
}
