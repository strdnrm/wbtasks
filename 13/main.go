package main

import "fmt"

func main() {
	a := 10
	b := 17
	fmt.Println("Before swap: ", a, b)
	swapInts(&a, &b)
	fmt.Println("After swap: ", a, b)
	a, b = b, a
	fmt.Println("After swap: ", a, b)
}

func swapInts(a, b *int) {
	*b += *a
	*a = *b - *a
	*b -= *a
}
