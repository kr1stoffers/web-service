package main

// Сдвиг массива: Реализуйте циклический сдвиг элементов массива влево на
// одну позицию (первый элемент становится последним).

import "fmt"

func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("Исходный массив: %v\n", arr)

	first := arr[0]
	for i := range len(arr) - 1 {
		arr[i] = arr[i+1]
	}
	arr[len(arr)-1] = first

	fmt.Printf("Сдвинутый массив: %v\n", arr)
}
