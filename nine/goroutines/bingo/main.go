package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

func main() {
	sum := 0
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				mu.Lock()
				sum++
				mu.Unlock()
			}

		}()
	}
	wg.Wait()
	fmt.Println(sum)
}
