/*
9. Генератор случайных строк: Сгенерируйте случайную строку заданной длины
из букв латинского алфавита.
*/
package main

import (
	"fmt"
	"math/rand"
)

func generateRandomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]rune, length)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func main() {
	fmt.Println(generateRandomString(15))
}
