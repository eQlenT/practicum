package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func action(ctx context.Context) {
	defer wg.Done()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println(ctx.Value("func"), ctx.Value("uid"), ctx.Value("user"))
	case <-ctx.Done():
		fmt.Println("Context canceled")
		return
	}
}

func main() {
	ctx := context.WithValue(context.Background(), "func", "main")
	ctx1 := context.WithValue(ctx, "uid", 701)
	wg.Add(1)
	go action(ctx1)
	// Теперь мы не вызываем ctx1.Done() здесь, так как мы хотим, чтобы функция action выполнилась
	// Ожидаем завершения action
	// После того как action завершилась, мы можем отменить контекст
	time.Sleep(300 * time.Millisecond)
	ctx1.Done()
	wg.Wait()
}
