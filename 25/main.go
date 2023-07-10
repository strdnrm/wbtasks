package main

import (
	"flag"
	"fmt"
	"time"
)

// Позволяет приостановить выполнение программы на заданное количество секунд
func Sleep(seconds int) {
	timer := time.NewTimer(time.Duration(seconds) * time.Second)
	<-timer.C
}

func main() {
	var n int
	flag.IntVar(&n, "n", 3, "sleep seconds")
	flag.Parse()
	Sleep(n) // Приостанавливаем выполнение программы на n секунд
	fmt.Printf("Work completed in %d seconds\n", n)
}
