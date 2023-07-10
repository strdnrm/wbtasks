package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	twopowtwenty := new(big.Int).Exp(big.NewInt(2), big.NewInt(20), nil)
	a := big.NewInt(twopowtwenty.Int64())
	b := big.NewInt(twopowtwenty.Int64() * 2)
	c := big.NewInt(0)
	fmt.Printf("a: %d b: %d\n", a, b)
	fmt.Println("a / b =", c.Div(a, b))
	fmt.Println("a * b =", c.Mul(a, b))
	fmt.Println("a + b =", c.Add(a, b))
	fmt.Println("a - b =", c.Sub(a, b))
}
