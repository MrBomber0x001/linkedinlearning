package main

import (
	"log"
	"os"
)

func CheckFileExists(filepath string) bool {
	if _, err := os.Stat(filepath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
func WriteFileData() {
	f, err := os.Create("text.txt")
	if err != nil {
		log.Fatal("Error occured", err.Error())
	}

	f.WriteString("welcome to my new file\n")

	data1 := []byte{'a', 'n', 'c', '\n'}

	f.Write(data1)
	defer f.Close()
	f.Sync() // forces the data to be written on the disk!
}

func appendFileData(fname string, data string) {
	// use the low-level OpenFile function to open the file in append file mode!
	f, err := os.OpenFile(fname, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error occurred", err.Error())
	}

	defer f.Close()

	_, err = f.WriteString(data)
}
func main() {
	filePath := "sample.txt"
	if !CheckFileExists(filePath) {
		WriteFileData()
	} else {
		os.Truncate(filePath, 1) // truncate the file to the fisrt 10 bytes!
	}

	appendFileData("sample.txt", "This would be added! and this also, besides this")
}
