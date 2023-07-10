package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Before delete", numbers)

	// Удаляем i-ый элемент из слайса
	i := 2
	numbers = append(numbers[:i], numbers[i+1:]...)

	fmt.Println("After delete: ", numbers)
}
