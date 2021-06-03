package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Book struct {
	Id     int64
	Author string
	Title  string
}

func init() {
	// register model
	orm.RegisterModel(new(Book))
}
func (b Book) String() string {
	return fmt.Sprintf("Book{Id:%v Author: %s Title: %s}", b.Id, b.Author, b.Title)
}

// func init() {
// 	// register model
// 	orm.RegisterModel(new(Book))

// 	// set default database
// 	orm.RegisterDataBase("default", "mysql", "root:root@/my_db?charset=utf8")
// }
