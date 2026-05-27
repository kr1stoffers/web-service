package main

// Приветствие: Напишите программу, которая запрашивает имя пользователя и
// выводит приветствие в формате "Привет, <имя>!".

import (
	"fmt"
)

func main() {
	var name string
	fmt.Print("Введите ваше имя: ")
	fmt.Scanln(&name)
	fmt.Printf("Привет, %s!\n", name)
}
