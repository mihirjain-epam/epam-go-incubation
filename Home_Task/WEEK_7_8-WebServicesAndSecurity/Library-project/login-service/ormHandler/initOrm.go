package ormHandler

import (
	"sync"

	"github.com/astaxie/beego/orm"
)

var instantiated orm.Ormer = nil
var once sync.Once

func New() *orm.Ormer {
	once.Do(func() {
		orm.RegisterDataBase("default", "mysql", "root:root@/my_db?charset=utf8")
		orm.RegisterDriver("mysql", orm.DRMySQL)
		instantiated = orm.NewOrm()
	})
	return &instantiated
}
