package panic

import "fmt"

func StartWebServer() {
	fmt.Println("Starting web server")
	if true {
		panic("Something wrong happened")
	}
	fmt.Println("Started web server")
}
