package main

import "fmt"

type sample struct {
	id   int
	name string
}

func main() {
	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])

	m["foo"] = 22
	m["tess"] = m["tess"] + 5
	fmt.Println(m["foo"])

	delete(m, "foo")
	fmt.Println(m["foo"])

	anotherMap := map[int]sample{1: sample{id: 1, name: "mihir"}}
	fmt.Println(anotherMap[1])

}
