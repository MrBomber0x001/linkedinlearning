package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	const httpbin = "https://httpbin.org/get"

	resp, err := http.Get(httpbin)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	fmt.Println("Status: ", resp.Status)
	fmt.Println("protocol: ", resp.Proto)
	fmt.Println("Content Length: ", resp.ContentLength)
	fmt.Println("Status Code: ", resp.StatusCode)

	// use a string builder to build a content from bytes
	var sb strings.Builder
	content, err := ioutil.ReadAll(resp.Body)
	bytesCount, err := sb.Write(content)

	fmt.Println(bytesCount, sb.String())

}
