package main

import (
	"regexp"
	"strings"
)

var (
	wordRe = regexp.MustCompile(`\w+`)
)

func Tokenize(text string) []string {
	var tokens []string
	for _, tok := range wordRe.FindAllString(text, -1) {
		tok = strings.ToLower(tok)
		tokens = append(tokens, tok)
	}

	return tokens
}
