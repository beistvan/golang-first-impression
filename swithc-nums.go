package main

import "fmt"

func main() {
	var n int
	s := ""
	for i := 0; i < 3; i++ {
		fmt.Scanln(&n)
		switch n {
		case 0:
			s = "Zero"
		case 1:
			s = "One"
		case 2:
			s = "Two"
		case 3:
			s = "Three"
		case 4:
			s = "Four"
		case 5:
			s = "Five"
		case 6:
			s = "Six"
		case 7:
			s = "Seven"
		case 8:
			s = "Eight"
		case 9:
			s = "Nine"
		case 10:
			s = "Ten"
		}
		fmt.Println(s)
	}

}
