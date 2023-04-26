package main

import (
	"encoding/xml"
	"fmt"
)

type person struct {
	XMLName    xml.Name `xml:"persondata"`
	Name       string   `xml:"fullname"`
	Address    string   `xml:"addre"`
	Age        int      `xml:"age,attr"`
	FaveColors []string `xml:"fave_colors"`
}

func encodeExample() {
	people := []person{
		{Name: "yousf meska", Address: "24 district", Age: 25, FaveColors: nil},
		{Name: "john public", Address: "456 bla bla bla", Age: 31, FaveColors: []string{"purple", "yello"}},
	}

	result, err := xml.MarshalIndent(people, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s%s\n", xml.Header, result)
}

func main() {
	encodeExample()
}
