package main

import (
	"flag"
	"fmt"
	"time"
)

func sender(dataChan chan<- int) {
	for i := 1; ; i++ {
		dataChan <- i
		time.Sleep(time.Millisecond * 300)
	}
}

func receiver(dataChan <-chan int, done chan<- bool) {
	for data := range dataChan {
		fmt.Println(data)
	}
	done <- true
}

func main() {
	var n int
	flag.IntVar(&n, "n", 5, "amount of work seconds")
	flag.Parse()

	timer := time.NewTimer(time.Duration(n) * time.Second)

	dataChan := make(chan int)
	done := make(chan bool)

	go sender(dataChan)         //Заруск отправителя в канал
	go receiver(dataChan, done) //Запуск получателя из канала

	select {
	case <-timer.C:
		close(dataChan)
		<-done
		fmt.Println("Work completed")
	case <-done:
		fmt.Println("Work completed")
	}
}
