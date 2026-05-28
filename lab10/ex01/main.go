/*
1. Эхо-сервер с завершением: Модифицируйте эхо-сервер так, чтобы он
завершал соединение, если клиент отправил "exit".
*/
package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.TrimSpace(strings.ToLower(text)) == "exit" {
			fmt.Println("До свидания!")
			fmt.Fprintln(conn, "До свидания!")
			break
		}
		fmt.Fprintln(conn, "Эхо:", text)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
		return
	}
	defer listener.Close()
	fmt.Println("TCP Эхо-сервер запущен на порту 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
	}
}
