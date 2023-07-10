package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(id int, dataChan <-chan string, wg *sync.WaitGroup) {
	for data := range dataChan {
		fmt.Printf("Worker %d: %s\n", id, data)
	}
	wg.Done()
}

func main() {
	var numWorkers int
	flag.IntVar(&numWorkers, "n", 5, "amount of workers")
	flag.Parse()

	fmt.Println("Number of workers: ", numWorkers)
	dataChan := make(chan string)

	wg := &sync.WaitGroup{}
	wg.Add(numWorkers)

	//Запуск воркеров для чтения из канала
	for i := 1; i <= numWorkers; i++ {
		go worker(i, dataChan, wg)
	}

	go func() {
		for {
			time.Sleep(time.Millisecond * 300)
			dataChan <- "test"
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalChan
		close(dataChan)
	}()

	wg.Wait()

	fmt.Println("Work completed")
}
