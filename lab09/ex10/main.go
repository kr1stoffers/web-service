/*
10. Дополнительно: Изучите пакет io/ioutil (устаревший) и его современные
аналоги os.ReadFile, os.WriteFile. Напишите пример использования.
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "example_ioutil.txt"
	dataToWrite := "Этот текст записан с помощью современных функций пакета os."

	// Начиная с Go 1.16 пакет io/ioutil объявлен УСТАРЕВШИМ (deprecated).
	// Вместо ioutil.WriteFile теперь нужно использовать os.WriteFile
	err := os.WriteFile(filename, []byte(dataToWrite), 0644)
	if err != nil {
		fmt.Println("Ошибка записи файла:", err)
		return
	}
	fmt.Println("Файл успешно записан.")

	// Вместо ioutil.ReadFile теперь нужно использовать os.ReadFile
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	fmt.Println("Содержимое прочитанного файла:")
	fmt.Println(string(content))

	// Удаляем временный файл за собой
	_ = os.Remove(filename)
}
