package main

import (
	"fmt"
	"time"
)

func main() {
	for _, v := range []string{".", "o"} {
		// для горутины используется анонимная функция
		go func(s string) {
			for i := 0; i < 30; i++ {
				fmt.Print(s)
				time.Sleep(time.Millisecond)
			}
		}(v) // передаём параметр в анонимную функцию
	}
	time.Sleep(1000 * time.Millisecond)
}
