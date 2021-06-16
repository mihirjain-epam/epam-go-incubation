package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Port     int
	BasePath string
	UsersURI string
	BooksURI string
}

var Config Configuration

func init() {
	file, err := os.Open("conf.json")
	if err != nil {
		fmt.Println("cannot open config file")
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		fmt.Println("cannot decode config file")
		panic(err)
	}
	Config.UsersURI = os.Getenv("UsersServiceURL") + Config.UsersURI
	Config.BooksURI = os.Getenv("BooksServiceURL") + Config.BooksURI
}
