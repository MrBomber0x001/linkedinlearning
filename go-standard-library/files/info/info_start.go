package main

import (
	"fmt"
	"os"
)

func checkFileExists(filepath string) bool {
	if _, err := os.Stat(filepath); err != nil {
		fmt.Println(filepath)
		fmt.Println(err)
		if os.IsNotExist(err) {
			fmt.Println("got here")
			return false
		}
	}
	return true
}
func main() {
	// stats, _ := os.Stat("notes.md")
	// fmt.Println("Modification time", stats.ModTime())
	// fmt.Println("File mode: ", stats.Mode())
	// fmode := stats.Mode()
	// if fmode.IsRegular() {
	// 	fmt.Println("Is regular file")
	// }

	// fmt.Println("file size", stats.Size())
	// fmt.Println("Is dir?", stats.IsDir())
	exists := checkFileExists("notes.md")
	fmt.Println("File exists check: ", exists)
}
