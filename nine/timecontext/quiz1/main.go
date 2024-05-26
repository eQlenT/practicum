package main

import (
	"fmt"
	"time"
)

func main() {
	// напишите код программы
	ticker := time.NewTicker(200 * time.Millisecond)
	var count int
	for {
		select {
		case <-ticker.C:
			if count == 20 {
				return
			}
			count++
			if count%5 == 0 {
				fmt.Print("o")
				break
			}
			fmt.Print(".")
		}
	}
}
