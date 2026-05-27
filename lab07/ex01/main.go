/*
1. Параллельный вывод: Запустите 5 горутин, каждая из которых выводит свой
номер и сообщение "start", затем "end". Используйте WaitGroup.
*/
package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Горутина %d: start\n", id)
	fmt.Printf("Горутина %d: end\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}
