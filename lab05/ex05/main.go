/*
5. Счётчик слов: Разбейте текст на слова и выведите топ-3 самых частых слова
(можно использовать карту).
*/
package main

import (
	"fmt"
	"sort"
	"strings"
)

type wordCount struct {
	word  string
	count int
}

func printTopThreeWords(text string) {
	words := strings.Fields(strings.ToLower(text))
	freq := make(map[string]int)
	for _, word := range words {
		word = strings.Trim(word, ".,!?;:")
		if word != "" {
			freq[word]++
		}
	}

	var list []wordCount
	for k, v := range freq {
		list = append(list, wordCount{k, v})
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].count > list[j].count
	})

	fmt.Println("Топ-3 частых слова:")
	for i := 0; i < 3 && i < len(list); i++ {
		fmt.Printf("%s: %d\n", list[i].word, list[i].count)
	}
}

func main() {
	text := "Разбейте текст на слова и выведите топ-3 самых частых слова. Обратный порядок слова. Плохой текст."
	printTopThreeWords(text)
}
