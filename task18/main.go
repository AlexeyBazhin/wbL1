package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := &Counter{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter.Inc()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter.GetValue())
}
