package main

import (
	"fmt"
	"unicode/utf8"
)

func reverseString(input string) string {
	length := len(input)
	runes := make([]rune, length)
	for _, runeValue := range input {
		length -= utf8.RuneLen(runeValue)
		runes[length] = runeValue
	}
	return string(runes)
}

func main() {
	input := "главрыба"
	reversed := reverseString(input)
	fmt.Println("Reversed string:", reversed)
}
