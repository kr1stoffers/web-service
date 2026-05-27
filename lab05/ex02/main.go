/*
2. Палиндром: Проверьте, является ли строка палиндромом (игнорируя пробелы
и знаки препинания). Учитывайте Unicode.
*/
package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	var cleanRunes []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cleanRunes = append(cleanRunes, unicode.ToLower(r))
		}
	}

	for i, j := 0, len(cleanRunes)-1; i < j; i, j = i+1, j-1 {
		if cleanRunes[i] != cleanRunes[j] {
			return false
		}
	}
	return true
}

func main() {
	text1 := "А роза упала на лапу Азора"
	text2 := "Привет, мир!"
	fmt.Println(text1, "->", isPalindrome(text1))
	fmt.Println(text2, "->", isPalindrome(text2))
}
