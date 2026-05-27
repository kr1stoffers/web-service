/*
3.  Конвейер чисел: Создайте конвейер из трёх горутин: первая генерирует числа
от 1 до 10, вторая умножает их на 2, третья выводит. Используйте два канала.
*/
package main

import "fmt"

func generator(out chan<- int) {
	for i := 1; i <= 10; i++ {
		out <- i
	}
	close(out)
}

func multiplier(in <-chan int, out chan<- int) {
	for num := range in {
		out <- num * 2
	}
	close(out)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go generator(ch1)
	go multiplier(ch1, ch2)

	for res := range ch2 {
		fmt.Println("Получено:", res)
	}
}
