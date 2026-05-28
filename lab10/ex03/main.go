/*
3. HTTP-сервер с маршрутами: Создайте HTTP-сервер с несколькими
маршрутами: /hello, /time (возвращает текущее время), /status
(возвращает "OK").
*/
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Привет")
	})

	mux.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Текущее время сервера:", time.Now().Format(time.RFC3339))
	})

	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	})

	fmt.Println("HTTP-сервер запущен на http://localhost:8082")
	http.ListenAndServe(":8082", mux)
}
