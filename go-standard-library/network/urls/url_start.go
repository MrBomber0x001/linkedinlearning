package main

import (
	"fmt"
	"net/url"
)

func main() {
	s := "https://yousefmeska.tech/user?username=mrbomber0x001"
	result, _ := url.Parse(s)
	fmt.Println(result.Scheme, result.Host, result.Path, result.User, result.RawQuery, result.Port())
	vals := result.Query()
	fmt.Println(vals["username"])
	// construct url
	newurl := &url.URL{
		Scheme:   "https",
		Host:     "www.example.com",
		Path:     "/args",
		RawQuery: "x=1&y=2",
	}
	ss := newurl.String()
	fmt.Println(ss)
	newurl.Host = "anotherone.com"
	ss = newurl.String()
	fmt.Println(ss)

	// create a new Values struct and encode it as a query string
	newVals := url.Values{}
	newVals.Add("x", "100")
	newurl.RawQuery = newVals.Encode()
	ss = newurl.String()
	fmt.Println(ss)
}
