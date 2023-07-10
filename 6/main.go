package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(stopChan <-chan bool, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		default:
			// Работа выполняется
			fmt.Println("Wroking...")
			time.Sleep(time.Second)
		case <-stopChan:
			// Получен сигнал остановки через канал
			fmt.Println("Stop signal from channel")
			return
		case <-ctx.Done():
			// Получен сигнал отмены через контекст
			fmt.Println("Stop signal from context")
		}
	}
}

func main() {
	stopChan := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	var wg sync.WaitGroup
	wg.Add(1)

	go worker(stopChan, ctx, &wg)

	time.Sleep(3 * time.Second) // Подождать некоторое время

	stopChan <- true // Отправить сигнал остановки через канал
	cancel()         // Отменить выполнение горутины через контекст

	wg.Wait() // Ожидать завершения горутины
	fmt.Println("Work completed")
}
