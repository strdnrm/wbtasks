package main

import (
	"fmt"
	"strings"
)

func reverseWords(input string) string {
	words := strings.Split(input, " ")
	reversedWords := make([]string, len(words))
	for i, word := range words {
		reversedWords[len(words)-1-i] = word
	}
	return strings.Join(reversedWords, " ")
}

func main() {
	input := "snow dog sun"
	reversed := reverseWords(input)
	fmt.Println("Reversed words:", reversed)
}
