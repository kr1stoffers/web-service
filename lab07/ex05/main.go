/*
5. Таймер с select: Напишите программу, которая запускает горутину,
выполняющую долгую операцию (например, time.Sleep(5*time.Second)).
В главной горутине используйте select с таймаутом 2 секунды, чтобы
прервать ожидание, если операция не завершилась.
*/
package main

import (
	"fmt"
	"time"
)

func longOperation(ch chan<- bool) {
	time.Sleep(5 * time.Second)
	ch <- true
}

func main() {
	ch := make(chan bool)
	go longOperation(ch)

	select {
	case <-ch:
		fmt.Println("Операция успешно завершена!")
	case <-time.After(2 * time.Second):
		fmt.Println("Таймаут! Превышено время ожидания 2 секунды.")
	}
}
