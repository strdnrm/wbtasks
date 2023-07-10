package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	data := make(map[string]int)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			// Захватываем мьютекс перед записью в map
			mu.Lock()
			// Выполняем запись в map
			key := fmt.Sprintf("key%d", index)
			data[key] = index
			mu.Unlock()
			fmt.Printf("Written value %d for key %s\n", index, key)
			wg.Done()
		}(i)
	}

	wg.Wait()

	// Выводим содержимое map
	for key, value := range data {
		fmt.Printf("Key: %s, value: %d\n", key, value)
	}
}
