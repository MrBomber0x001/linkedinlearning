package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	postRequestTest()
}

func postRequestTest() {
	const httpbin = "https://httpbin.org/post"
	// TODO: POST operation using Post
	reqBody := strings.NewReader(`
	{
		"field1": "This is field 1",
		"field2": 210	
	}
	`)

	resp, err := http.Post(httpbin, "application/json", reqBody)
	if err != nil {
		return
	}
	content, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Printf("%s\n", content)

	// TODO: POST operation using PostForm

	data := url.Values{}
	data.Add("field1", "Field added via values")
	data.Add("field2", "300")
	resp, err = http.PostForm(httpbin, data)
	content, _ = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Printf("%s\n", content)
}
