package main

import (
	"encoding/json"
	"fmt"
)

type myStruct struct {
	uuid string
	Name string
}

func (m myStruct) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Uuid string
		Name string
	}{
		Uuid: m.uuid,
		Name: m.Name,
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func main() {
	myStruct := myStruct{uuid: "PROPER-UUID-STRING", Name: "Some Proper Name"}
	j, err := json.Marshal(myStruct)
	if err != nil {
		fmt.Println("Something went wrong")
		return
	}
	fmt.Println(string(j))
}
