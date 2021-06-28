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

func (u UserDTO) String() string {
	return fmt.Sprintf("UserDTO {Id:%v Name: %s BirtDate: %v Associated books:%v}", u.Id, u.Name, u.BirthDate, u.AssociatedBooks)
}
