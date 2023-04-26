package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	tempPath := os.TempDir()
	fmt.Println(tempPath)

	tempContent := []byte("This is some temp content for the file")
	tmpFile, err := ioutil.TempFile(tempPath, "tempfile_*.txt")
	if err != nil {
		panic(err)
	}

	if _, err := tmpFile.Write(tempContent); err != nil {
		panic(err)
	}

	data, _ := ioutil.ReadFile(tmpFile.Name())
	fmt.Printf("%s\n", data)

	fmt.Println("Temp file is named:", tmpFile.Name())
	defer os.Remove(tmpFile.Name())

	tmpDir, err := ioutil.TempDir("", "tempdir*")
	if err != nil {
		panic(err)
	}
	fmt.Println("the temp dir is named:", tmpDir)
	defer os.RemoveAll(tmpDir)
}
