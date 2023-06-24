package main

import "fmt"

func changeStrings(a *string, b *string) {
	*a, *b = *b, *a
}

func main() {

	hello, bello := "Hello", "Bello"

	fmt.Println(hello, bello)

	changeStrings(&hello, &bello)

	fmt.Println(hello, bello)
}
