package main

import "fmt"

func main() {

	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
	var n string
	fmt.Scanln(&n)
	results = append(results, n)
	s := 0
	for _, v := range results {
		switch v {
		case "w":
			s += 3
		case "d":
			s += 1
		}
	}
	fmt.Println(s)
}
