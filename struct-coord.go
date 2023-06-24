package main

import "fmt"

type Coord struct {
	x int
	y int
}

func main() {
	a := Coord{7, 5}
	fmt.Println(a.x - a.y)
}
