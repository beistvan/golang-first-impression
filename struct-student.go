
package main

import "fmt"

type Student struct {
    name, university string    
}

func main() {
	student := Student{"John Smith", "MIT"}
	fmt.Println(student.name)
	fmt.Println(student.university)
}

