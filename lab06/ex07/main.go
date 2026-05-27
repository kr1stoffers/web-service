/*
7. Пустой интерфейс в карте: Создайте карту map[string]interface{} и
заполните её значениями разных типов. Выведите все ключи и значения.
*/
package main

import "fmt"

func main() {
	data := map[string]interface{}{
		"name":    "Алексей",
		"age":     21,
		"gpa":     4.75,
		"isReady": true,
	}

	for key, value := range data {
		fmt.Printf("Ключ: %-8s | Тип: %-10T | Значение: %v\n", key, value, value)
	}
}
