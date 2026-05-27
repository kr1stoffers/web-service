/*
1. Частота символов: Напишите функцию, которая принимает строку и
возвращает карту с частотой каждого символа (руны).
*/
package main

import "fmt"

func countCharFrequency(s string) map[rune]int {
	frequency := make(map[rune]int)
	for _, char := range s {
		frequency[char]++
	}
	return frequency
}

func main() {
	text := "привет, мир"
	freq := countCharFrequency(text)
	for char, count := range freq {
		fmt.Printf("%c: %d\n", char, count)
	}
}
