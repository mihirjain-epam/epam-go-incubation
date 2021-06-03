package main

import (
	"fmt"
	"log"
	"net/http"

	"epam.com/web-services/library-management/users-service/handlers"
	"github.com/astaxie/beego/orm"
)

func main() {
	handlers.SetupRoutes()
	err := http.ListenAndServe(":5001", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// orm functionality to drop and create `users` table
func dropAndCreateUserTable() {
	name := "default"

	force := true

	verbose := true

	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}
