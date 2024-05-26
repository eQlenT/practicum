package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		m  sync.Map
		wg sync.WaitGroup
	)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		// все горутины делают одну и ту же работу
		go func() {
			defer wg.Done()
			for j := 1; j <= 10; j++ {
				// пишем данные
				m.Store(j, j*j)
			}
		}()
	}
	wg.Wait()
	sum := 0
	m.Range(func(key, value any) bool {
		if v, ok := value.(int); ok {
			sum += v
		}
		return true
	})
	fmt.Println(sum)
}
