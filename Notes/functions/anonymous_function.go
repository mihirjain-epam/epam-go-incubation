package main

import "fmt"

func f1() {
	value := func(s string) {
		fmt.Println("Welcome! to " + s)
	}
	value("Mihir")
}

func f2() func(i, j string) string {
	myf := func(i, j string) string {
		return i + j
	}
	return myf
}

func f3(f func(i, j string) string, s string) string {
	return f("Hello ", "World ") + s
}

func main() {
	f1()
	an1 := f2()
	fmt.Println(an1("Hello ", "World "))
	fmt.Println(f3(f2(), "Mihir"))

}
