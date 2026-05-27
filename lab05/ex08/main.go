/*
8. Сложение больших чисел: Напишите функцию, которая принимает две
строки, содержащие целые числа, и возвращает их сумму в виде строки (числа
могут быть очень большими, используйте сложение столбиком).
*/
package main

import "fmt"

func addLargeNumbers(num1, num2 string) string {
	runes1 := []rune(num1)
	runes2 := []rune(num2)

	var result []rune
	carry := 0
	i, j := len(runes1)-1, len(runes2)-1

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(runes1[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(runes2[j] - '0')
			j--
		}
		carry = sum / 10
		result = append([]rune{rune((sum % 10) + '0')}, result...)
	}

	return string(result)
}

func main() {
	n1 := "999999999999999999999999999999"
	n2 := "1"
	fmt.Println(addLargeNumbers(n1, n2))
}
