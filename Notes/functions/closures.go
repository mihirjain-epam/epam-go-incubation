package main

import "fmt"

// closures use property of anonymous functions to enclose data inside the calling function/scope

func newCounter() func() int {
	counter := 0
	return func() int {
		counter += 1
		return counter
	}
}

func main() {
	count := newCounter()
	fmt.Println(count())
	fmt.Println(count())
}
