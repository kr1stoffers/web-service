/*
4.  Пул воркеров: Реализуйте пул из 3 воркеров, которые читают задания из
канала (например, числа) и возвращают результат (квадрат) в другой канал.
Главная горутина отправляет 10 заданий и закрывает канал.
*/
package main

import (
	"fmt"
	"sync"
)

func workerPool(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		results <- job * job
	}
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go workerPool(w, jobs, results, &wg)
	}

	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
	close(results)

	for res := range results {
		fmt.Println("Результат из пула:", res)
	}
}
