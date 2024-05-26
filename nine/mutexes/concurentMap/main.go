package main

import (
	"fmt"
	"time"
)

type Cache struct {
	m map[int]int
}

func (cache *Cache) Get(i int) int {
	v, ok := cache.m[i]
	if ok {
		return v
	}
	// получаем значение для указанного ключа
	v = 2 * i
	cache.m[i] = v
	return v
}

func main() {
	cache := Cache{
		m: make(map[int]int),
	}
	for i := 0; i < 20; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				cache.Get(j)
			}
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(len(cache.m))
}
