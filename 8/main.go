package main

import (
	"flag"
	"fmt"
)

func main() {
	var num int64 = 10
	var i int
	flag.IntVar(&i, "n", 2, "bit number to be set")
	flag.Parse()

	// Устанавливаем i-й бит в 1
	num |= 1 << i
	fmt.Printf("Number after setting the %dth bit to 1: %d\n", i, num)

	// Устанавливаем i-й бит в 0
	num &= ^(1 << i)
	fmt.Printf("Number after setting the %dth bit to 0: %d\n", i, num)
}
