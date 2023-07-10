package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}

	input := make(chan int)
	output := make(chan int)

	// Горутина для записи чисел из массива в канал input
	go func() {
		for _, num := range nums {
			input <- num
		}
		close(input)
	}()

	// Горутина для умножения чисел из канала input на 2 и записи результата в канал output
	go func() {
		for num := range input {
			output <- num * 2
		}
		close(output)
	}()

	// Вывод чисел из канала output в stdout
	for num := range output {
		fmt.Println(num)
	}
}
