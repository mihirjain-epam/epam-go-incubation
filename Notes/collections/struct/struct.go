package main

import "fmt"

type user struct {
	ID        int
	FirstName string
	LastName  string
}

func main() {
	var u user
	u.ID = 1
	u.FirstName = "Mihir"
	u.LastName = "Jain"
	fmt.Println(u)

	u2 := user{ID: 1,
		FirstName: "Manit",
		LastName:  "Jain",
	}
	fmt.Println(u2)
}
