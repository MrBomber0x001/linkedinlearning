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

func main() {
	xmldata := `
	<persondata age="31">
		<fullname>john public</fullname>
		<addre>456 bla bla bla</addre>
		<fave_colors>purple</fave_colors>
		<fave_colors>yello</fave_colors>
	</persondata>
	`

	var p person

	xml.Unmarshal([]byte(xmldata), &p)
	fmt.Printf("%#v\n", p)
}
