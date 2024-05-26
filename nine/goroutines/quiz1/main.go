package main

import (
	"fmt"
	"time"
)

const Count = 1000
const Lines = 4

var data [Lines][Count]int

func main() {
	// вставьте сюда код с запуском четырёх горутин
	go fullfil(0)
	go fullfil(1)
	go fullfil(2)
	go fullfil(3)

	time.Sleep(2000 * time.Millisecond)
	// проверим как заполнен массив
	sum := 0
	for i := 0; i < Lines; i++ {
		for j := 0; j < Count; j++ {
			sum += data[i][j]
		}
	}
	fmt.Println(sum)
}

func fullfil(line int) {
	for i := 0; i < 1000; i++ {
		data[line][i] = i
	}
}
