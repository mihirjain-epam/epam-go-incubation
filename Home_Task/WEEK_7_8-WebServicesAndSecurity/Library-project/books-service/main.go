package main

import (
	"fmt"
	"log"
	"net/http"

	"epam.com/web-services/library-management/books-service/config"
	"epam.com/web-services/library-management/books-service/handlers"
)

func main() {
	handlers.SetupRoutes()
	err := http.ListenAndServe(":"+fmt.Sprint(config.Config.Port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
