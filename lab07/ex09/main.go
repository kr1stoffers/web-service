/*
9.  Синхронизация через канал: Используйте канал для синхронизации запуска
нескольких горутин: главная горутина отправляет сигнал всем дочерним, и
они начинают работу.
*/
package main

import (
	"fmt"
	"sync"
)

func workerReady(id int, startSignal <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Воркер %d готов и ждет сигнала...\n", id)
	<-startSignal
	fmt.Printf("Воркер %d начал выполнение работы!\n", id)
}

func main() {
	var wg sync.WaitGroup
	startSignal := make(chan struct{})

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go workerReady(i, startSignal, &wg)
	}

	fmt.Println("Главная горутина подготавливает данные...")
	fmt.Println("ВНИМАНИЕ... МАРШ!")
	close(startSignal) // Закрытие канала мгновенно разблокирует все горутины, читающие из него

	wg.Wait()
}
