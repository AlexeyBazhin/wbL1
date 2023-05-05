package main

import (
	"fmt"
	"sync"
)

func main() {
	counters := make(map[int]int)
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			for j := 0; j < 5; j++ {

				mu.Lock() //также можно использовать atomic
				counters[i] += j
				mu.Unlock()

			}
			wg.Done()
		}(i, wg)
	}
	wg.Wait()
	fmt.Println(counters)
}
