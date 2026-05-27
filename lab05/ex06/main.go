/*
6. Преобразование регистра: Создайте функцию, которая переводит первую
букву каждого слова в заглавную, остальные – в строчные.
*/
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func titleCase(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		runes := []rune(strings.ToLower(word))
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
		}
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
}

func main() {
	text := "пРиВЕТ оТлИЧНЫЙ языК gOlAnG"
	fmt.Println(titleCase(text))
}
