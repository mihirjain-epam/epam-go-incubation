package main

import (
	"fmt"

	"epam.com/go/training/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Mihir",
		LastName:  "Jain",
	}
	fmt.Println(u)
}
