/*
3. Обратный порядок слов: Напишите функцию, которая принимает
предложение и возвращает его с обратным порядком слов.
*/
package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	sentence := "раз два три"
	fmt.Println(reverseWords(sentence))
}
