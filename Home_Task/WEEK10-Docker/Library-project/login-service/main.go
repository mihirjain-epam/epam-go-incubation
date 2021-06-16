package main

import (
	"log"
	"net/http"

	"epam.com/web-services/library-management/login-service/handlers"
)

func main() {
	handlers.SetupRoutes()
	err := http.ListenAndServe(":5003", nil)
	if err != nil {
		log.Fatal(err)
	}
}
