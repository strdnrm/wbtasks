package main

import (
	"fmt"
	"sync"
)

// ...
func main() {
	nums := []int{2, 4, 6, 8, 10}
	wg := &sync.WaitGroup{}

	for _, num := range nums {
		wg.Add(1)
		//Вычиление квадратов
		go func(num int, wg *sync.WaitGroup) {
			square := num * num
			fmt.Println(square)
			wg.Done()
		}(num, wg)
	}

	wg.Wait()
}
