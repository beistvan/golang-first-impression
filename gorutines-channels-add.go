package main

import "fmt"

func download(n int, ch chan int) {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	ch <- sum
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	var s1, s2, s3 int
	fmt.Scanln(&s1)
	fmt.Scanln(&s2)
	fmt.Scanln(&s3)

	// s1, s2, s3 := 10, 100, 42 // sum = 6008

	go download(s1, ch1)
	go download(s2, ch2)
	go download(s3, ch3)

	// sum of all results
	fmt.Println(<-ch1 + <-ch2 + <-ch3)
}
