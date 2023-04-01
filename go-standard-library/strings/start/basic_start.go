package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "the quick brown fox jumps over the lazy dog"
	fmt.Println(len(s))

	for _, ch := range s {
		fmt.Print(string(ch), ",")
	}
	fmt.Println("")

	fmt.Println("dog" < "cat")

	// compare: [0] => are the same, [-1] if A is less than B, [+1] if A is bigger then B [not comparing by length]
	result := strings.Compare("dog", "cat") // not as efficent as using math operators!
	fmt.Println(result)
	result = strings.Compare("dog", "dog")
	fmt.Println(result)

	fmt.Println(strings.EqualFold("Hello", "hello"))

	s1 := strings.ToLower(s)
	s2 := strings.ToUpper(s)
	s3 := strings.Title(s)
	fmt.Println(s1, s2, s3)
}

func capitalize(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + word[1:]
		}
	}
	return strings.Join(words, "")
}
