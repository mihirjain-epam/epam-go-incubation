package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

// var block
var (
	user   []*User
	nextID = 1
)
