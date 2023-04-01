package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // as a stream
	s, _ := reader.ReadString('\n')     // until a fist occurance of delimiter
	fmt.Println(s)
}
