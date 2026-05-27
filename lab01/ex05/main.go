package main

// Поиск в массиве: Создайте массив из 10 целых чисел (можно задать при объявлении).
// Найдите минимальный и максимальный элементы, а также их индексы.

import "fmt"

func main() {
	arr := [10]int{4, 12, -3, 8, 15, 0, 23, -7, 11, 5}

	minVal, maxVal := arr[0], arr[0]
	minIdx, maxIdx := 0, 0

	for i := 1; i < len(arr); i++ {
		if arr[i] < minVal {
			minVal = arr[i]
			minIdx = i
		}
		if arr[i] > maxVal {
			maxVal = arr[i]
			maxIdx = i
		}
	}

	fmt.Printf("Массив: %v\n", arr)
	fmt.Printf("Минимум: %d (индекс %d)\n", minVal, minIdx)
	fmt.Printf("Максимум: %d (индекс %d)\n", maxVal, maxIdx)
}
