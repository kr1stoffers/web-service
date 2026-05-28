package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	targetURL := "https://httpbin.org"

	resp, err := http.Get(targetURL)
	if err != nil {
		fmt.Println("Ошибка HTTP-запроса:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return
	}

	fmt.Println("Ответ от тестового сервера получен успешно!")
	fmt.Println(string(body))
}
