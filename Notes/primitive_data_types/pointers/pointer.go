package main

import "fmt"

func main() {
	var firstName *string = new(string)
	*firstName = "mihir"
	fmt.Println(*firstName)
	lastName := "jain"
	ptr := &lastName
	fmt.Println(ptr, *ptr)

	lastName = "jai"
	fmt.Println(ptr, *ptr)
}
