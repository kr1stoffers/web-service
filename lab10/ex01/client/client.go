package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	// 1. Подключаемся к нашему Go-серверу
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Не удалось подключиться к серверу:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Успешно подключились к эхо-серверу!")

	reader := bufio.NewReader(conn)

	// Список сообщений, которые наш "робот" отправит серверу
	messages := []string{
		"Привет, сервер!",
		"Тестирую русский язык",
		"exit", // Финальная команда
	}

	for _, msg := range messages {
		// Делаем небольшую паузу между сообщениями, чтобы имитировать человека
		time.Sleep(1 * time.Second)

		// 2. Отправляем сообщение (обязательно добавляем \n в конец)
		fmt.Fprintf(conn, msg+"\n")
		fmt.Printf("Отправлено: %s\n", msg)

		// 3. Читаем, что ответил сервер
		reply, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Сервер закрыл соединение:", err)
			return
		}
		fmt.Print("Получен ответ: ", reply)
	}
}
