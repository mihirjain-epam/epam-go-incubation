package models

import (
	"fmt"

	"time"
)

type UserDTO struct {
	Id              int64
	Name            string
	BirthDate       time.Time
	AssociatedBooks []int64
}

func init() {
	// register model
	// orm.RegisterModel(new(UserDTO))
}
func (u UserDTO) String() string {
	return fmt.Sprintf("Book{Id:%v Author: %s Title: %v Associated books:%v}", u.Id, u.Name, u.BirthDate, u.AssociatedBooks)
}
