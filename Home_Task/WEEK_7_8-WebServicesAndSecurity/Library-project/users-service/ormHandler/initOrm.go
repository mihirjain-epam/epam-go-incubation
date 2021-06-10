package ormHandler

import (
	"sync"

	"epam.com/web-services/library-management/users-service/config"
	"github.com/astaxie/beego/orm"
)

var instantiated orm.Ormer = nil
var once sync.Once

func New() *orm.Ormer {
	once.Do(func() {
		orm.RegisterDataBase("default", config.Config.DBDriverName,
			config.Config.DBUserName+":"+
				config.Config.DBPassword+"@/"+
				config.Config.DBName+"?charset=utf8")
		orm.RegisterDriver("mysql", orm.DRMySQL)
		instantiated = orm.NewOrm()
	})
	return &instantiated
}
