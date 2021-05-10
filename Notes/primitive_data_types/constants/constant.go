package main

import "fmt"

const (
	// first = 1
	// second = "second"
	// iota has a rule - each time it is used it increments its value
	// first  = iota
	// second = iota

	first  = iota + 6
	second // iota +6
)

const (
	third = iota
)

func main() {
	// const c = 3
	// fmt.Println(c + 3)
	// fmt.Println(c + 1.2)

	fmt.Println(first, second, third)

}
