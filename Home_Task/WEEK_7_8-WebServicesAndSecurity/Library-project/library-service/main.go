package main

import (
	"fmt"
	"log"
	"net/http"

	"epam.com/web-services/library-management/library-service/config"
	"epam.com/web-services/library-management/library-service/handlers"
)

func main() {
	err := http.ListenAndServe(":"+fmt.Sprint(config.Config.Port), http.HandlerFunc(handlers.Serve))
	if err != nil {
		log.Fatal(err)
	}
}
