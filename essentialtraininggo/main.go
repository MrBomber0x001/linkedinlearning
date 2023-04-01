package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type Job struct {
	User   string `json:"user"`
	Action string `json:"action"`
	Count  int    `json:"count"`
}

func main() {
	// GET
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)

	// POST
	job := &Job{
		User:   "Meska",
		Action: "Punch",
		Count:  1,
	}
	var buff bytes.Buffer         // in memory reader and writer
	enc := json.NewEncoder(&buff) // encode the data into the buffer!
	if err := enc.Encode(job); err != nil {
		log.Fatal(err.Error())
	}

	resp, err = http.Post("https://httpbin.org/post", "application/json", &buff)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
