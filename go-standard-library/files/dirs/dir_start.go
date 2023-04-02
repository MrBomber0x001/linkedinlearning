package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	os.Mkdir("new", os.ModePerm)
	os.MkdirAll("path/to/new/dir", os.ModePerm)

	os.Remove("new")
	os.RemoveAll("path") // remove it and all its children

	dir, _ := os.Getwd()
	fmt.Println(dir)

	execdir, _ := os.Executable() // get the directory of the current process [executable made by the compiler at temp dir somewhere!]
	fmt.Println(execdir)

	// read the content of a directory
	contents, _ := ioutil.ReadDir(".")

	for _, fi := range contents {
		fmt.Println(fi.IsDir(), fi.Name())
	}
}
