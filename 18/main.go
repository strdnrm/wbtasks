package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	counter := &Counter{}

	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			counter.Inc()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("Value of counter: %d\n", counter.GetValue())
}
