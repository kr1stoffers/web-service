/*
10. Дополнительно: Изучите пакет unicode и напишите функцию, которая
удаляет из строки все символы, не являющиеся буквами или цифрами.
*/
package main

import (
	"fmt"
	"unicode"
)

func removeSpecialChars(s string) string {
	var result []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result = append(result, r)
		}
	}
	return string(result)
}

func main() {
	text := "Привет, мир! 123 @#$ gg..."
	fmt.Println(removeSpecialChars(text))
}
