package config

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Port         int
	BasePath     string
	UsersURI     string
	BooksURI     string
	DBName       string
	DBDriverName string
	DBUserName   string
	DBPassword   string
}

var Config Configuration

func init() {
	file, err := os.Open("conf.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		panic(err)
	}
}
