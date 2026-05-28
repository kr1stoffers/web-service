/*
9. Размер файла: Напишите функцию, которая возвращает размер файла в
байтах, не читая его полностью (используйте os.Stat).
*/
package main

import (
	"fmt"
	"os"
)

func getFileSize(filename string) (int64, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}

func main() {
	size, err := getFileSize("numbers.txt")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Printf("Размер файла numbers.txt: %d байт\n", size)
}
