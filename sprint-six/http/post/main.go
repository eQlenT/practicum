package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	for {

		response, err := http.Post("http://localhost:8080/", "text/plain", nil)
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
