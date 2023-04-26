package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name       string   `json:"fullname"`
	Address    string   `json:"addre"`
	Age        int      `json:"age"`
	FaveColors []string `json:"favecolors"`
}

func main() {
	decodeExample()
}

func decodeExample() {
	data := []byte(`
		{
			"fullname" : "meska",
			"addr" :"26 district",
			"age": 25,
			"favecolors" : ["purple", "white"]
		}
	`)
	var p person

	valid := json.Valid(data)
	if valid {
		fmt.Println(valid)
		json.Unmarshal(data, &p)
		fmt.Printf("%#v\n", p)
	}
	fmt.Println(valid)

	// TODO: data can also be decoded into a map structure
	var m map[string]interface{}
	json.Unmarshal(data, &m)
	fmt.Printf("%#v\n", p)

	for k, v := range m {
		fmt.Printf("key (%v) (%T: %v)\b", k, v, v)
	}
}
