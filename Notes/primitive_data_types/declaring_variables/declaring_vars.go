package main

import "fmt"

func main() {
	var c int
	c = 42
	fmt.Println(c)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "mihir"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	co := complex(3, 4)
	fmt.Println(co)

	r, i := real(co), imag(co)
	fmt.Println(r, i)
}
