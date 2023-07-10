package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	squareChan := make(chan int)

	for _, num := range nums {
		//Вычисление площайдей и запись их в канал
		go func(num int) {
			square := num * num
			squareChan <- square
		}(num)
	}

	sum := 0
	// Чтение площадей из канала и вычисление суммы
	for range nums {
		square := <-squareChan
		sum += square
	}

	close(squareChan)

	fmt.Println("Сумма квадратов чисел:", sum)
}
