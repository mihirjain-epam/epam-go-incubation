package main

import "fmt"

func main() {
	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])

	m["foo"] = 22

	fmt.Println(m["foo"])

	delete(m, "foo")
	fmt.Println(m)
}
