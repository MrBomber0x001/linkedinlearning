package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func getFileLength(filepath string) int64 {
	f, err := os.Stat(filepath)
	if err != nil {
		log.Fatal(err.Error())
	}
	return f.Size()
}
func main() {
	content, err := ioutil.ReadFile("notes.md")
	if err != nil {
		log.Fatal(err.Error())
	}

	// ReadFile reads as bytes, entire file into memory
	fmt.Println(string(content))

	// read smaller chunks
	const BuffSize = 20
	f, _ := os.Open("notes.md")
	defer f.Close()

	b1 := make([]byte, BuffSize)

	// read until as we reach the end!
	for {
		n, err := f.Read(b1)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err.Error())
			}
			break
		}

		fmt.Println("Bytes read", n)
		fmt.Println("content in bytes", b1[:n]) // read until the number of bytes read so far!
		fmt.Println("content as string", string(b1))
	}

	// we can read file at specific offset!
	l := getFileLength("notes.md")
	b2 := make([]byte, 10)
	_, err3 := f.ReadAt(b2, l-int64(len(b2))) // read 10 bytes from the end of the file!
	if err3 != nil {
		log.Fatal(err3.Error())
	}
}
