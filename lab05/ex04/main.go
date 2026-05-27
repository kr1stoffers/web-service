/*
4. Цензура: Замените все вхождения "bad" на "***" в строке.
*/
package main

import (
	"fmt"
	"strings"
)

func censorText(s string) string {
	return strings.ReplaceAll(s, "bad", "***")
}

func main() {
	text := "This is a bad example of a bad word."
	fmt.Println(censorText(text))
}
