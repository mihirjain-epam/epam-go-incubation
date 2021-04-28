package main

import (
	"net/http"

	"epam.com/go/training/controllers"
)

func main() {
	controllers.RegistrationControllers()
	http.ListenAndServe(":3000", nil)
}
