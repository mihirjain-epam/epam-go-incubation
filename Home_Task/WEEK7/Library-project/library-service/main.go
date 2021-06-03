package main

import (
	"fmt"
	"log"
	"net/http"

	"epam.com/web-services/library-management/library-service/handlers"
	"github.com/astaxie/beego/orm"
)

func main() {
	err := http.ListenAndServe(":5002", http.HandlerFunc(handlers.Serve))
	if err != nil {
		log.Fatal(err)
	}
}

// orm functionality to drop and create `library` table
func dropAndCreateLibraryTable() {
	name := "default"

	force := true

	verbose := true

	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}
