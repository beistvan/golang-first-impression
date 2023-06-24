package main

import "fmt"

func myStr(x, y string) (string, string) {
	return x, y
}

func main() {
	fmt.Println(myStr("Hello", "World!"))
}
