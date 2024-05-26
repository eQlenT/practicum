package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

/*
	 func main() {
		for {
			response, err := http.Get("http://localhost:8080/time")
			if err != nil {
				fmt.Println(err)
				return
			}
			body, err := io.ReadAll(response.Body)
			response.Body.Close()
			if err != nil {
				fmt.Println("Ошибка чтения:", err)
				break
			}
			fmt.Println(string(body))

			time.Sleep(5 * time.Second)
		}
	}
*/
func main() {
	client := &http.Client{
		Timeout: 1 * time.Second,
	}
	request, err := http.NewRequest(http.MethodGet, "http://localhost:8080/time", nil)
	if err != nil {
		fmt.Println("Ошибка формирования запроса:", err)
		return
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Ошибка отправки запроса:", err)
		return
	}
	// вызов после defer произойдёт при выходе из функции,
	// поэтому Close() с defer можно указать здесь, так как
	// response.Body.Close() нужно вызвать в любом случае
	defer response.Body.Close()
	// читаем тело ответа
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return
	}
	fmt.Println(string(body))
}
