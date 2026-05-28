package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
	"time"
)

func simulateUser(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		fmt.Printf("[%s] Ошибка подключения\n", name)
		return
	}
	defer conn.Close()

	// Горутина для фонового чтения сообщений от других участников чата
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Printf("--> %s услышал: %s\n", name, scanner.Text())
		}
	}()

	// Имитируем отправку сообщений
	time.Sleep(1 * time.Second)
	fmt.Fprintf(conn, "Всем привет, я "+name+"\n")

	time.Sleep(2 * time.Second)
	fmt.Fprintf(conn, "Кто-нибудь тут есть?\n")

	time.Sleep(2 * time.Second)
}

func main() {
	var wg sync.WaitGroup

	// Запускаем 3 виртуальных пользователей параллельно через WaitGroup
	users := []string{"Алиса", "Боб", "Чарли"}

	for _, user := range users {
		wg.Add(1)
		go simulateUser(user, &wg)
		time.Sleep(200 * time.Millisecond) // небольшая задержка перед входом следующего
	}

	wg.Wait()
	fmt.Println("Имитация чата завершена.")
}
