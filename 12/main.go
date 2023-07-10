package main

import "fmt"

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})

	for _, str := range sequence {
		set[str] = struct{}{}
	}

	fmt.Println(set)
}
