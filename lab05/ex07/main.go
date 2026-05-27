/*
7.  Удаление дубликатов: Удалите из строки все повторяющиеся символы
(оставить только первое вхождение).
*/
package main

import "fmt"

func removeDuplicateChars(s string) string {
	seen := make(map[rune]bool)
	var result []rune

	for _, char := range s {
		if !seen[char] {
			seen[char] = true
			result = append(result, char)
		}
	}
	return string(result)
}

func main() {
	text := "гооолаааннгг"
	fmt.Println(removeDuplicateChars(text))
}
