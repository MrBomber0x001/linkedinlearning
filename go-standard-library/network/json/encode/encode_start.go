package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name       string   `json:"fullname"`
	Address    string   `json:"addre"`
	Age        int      `json:"age"`
	FaveColors []string `json:"fave_color,omitempty"`
}

func encodeExample() {
	people := []person{
		{"yousf meska", "24 district", 25, nil},
		{"john public", "456 bla bla bla", 31, []string{"purple", "yello"}},
	}

	result, err := json.MarshalIndent(people, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", result)
}

func main() {
	encodeExample()
}
