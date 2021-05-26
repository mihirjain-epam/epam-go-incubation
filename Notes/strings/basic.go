package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	// var b byte = 120 //here allowed because 120 is an int
	// var s string = string(b)
	// var c byte = '1' // doesn't allow byte('120')  mostly because '120' is three bytes trying to assign in one
	// a := []byte{b}
	// fmt.Println(a)
	// fmt.Println(s)
	// fmt.Println(int(b))
	// fmt.Println(string(c))
	// fmt.Printf("%q", b)
	// fmt.Sprint(s)

	a := "plain string"
	first := a[0]
	fmt.Println(first)
	fmt.Println(reflect.TypeOf(first))

	str := "something"
	r := []rune(str)
	fmt.Printf("%q", r)

	copied := deepCopy(str)
	fmt.Println(&copied, &str)
	fmt.Println(&r[0], &str)

}

func deepCopy(s string) string {
	var sb strings.Builder
	sb.WriteString(s)
	return sb.String()
}
