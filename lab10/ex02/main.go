/*
2. Многопользовательский чат: Реализуйте простой TCP-чат: сервер принимает
подключения и рассылает сообщения от одного клиента всем остальным.
*/
package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

var (
	clients   = make(map[net.Conn]bool)
	clientsMu sync.Mutex
)

func broadcast(message string, sender net.Conn) {
	clientsMu.Lock()
	defer clientsMu.Unlock()
	for client := range clients {
		if client != sender {
			fmt.Fprint(client, message)
		}
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	clientsMu.Lock()
	clients[conn] = true
	clientsMu.Unlock()

	broadcast(fmt.Sprintf("[Система]: Подключился новый участник %s\n", conn.RemoteAddr()), conn)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := fmt.Sprintf("[%s]: %s\n", conn.RemoteAddr(), scanner.Text())
		broadcast(msg, conn)
	}

	clientsMu.Lock()
	delete(clients, conn)
	clientsMu.Unlock()

	broadcast(fmt.Sprintf("[Система]: Участник %s покинул чат\n", conn.RemoteAddr()), nil)
}

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Чат-сервер запущен на порту 8081...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}
