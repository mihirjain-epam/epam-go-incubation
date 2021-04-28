package conditionals

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func Testing_if() {
	u1 := User{
		ID:        1,
		FirstName: "Mihir",
		LastName:  "Jain",
	}
	u2 := User{
		ID:        2,
		FirstName: "Manit",
		LastName:  "Jain",
	}
	if u1.ID == u2.ID {
		fmt.Println("Same user!")
	} else if u1.FirstName == u2.FirstName {
		fmt.Println("Same user!")
	} else {
		fmt.Println("Different users!")
	}

}
