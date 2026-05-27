/*
10. Дополнительно: Изучите пакет atomic и замените мьютекс на атомарные
операции для инкремента счётчика.
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var counter int64 // Обязательно используем строго фиксированный int64

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1) // Безопасный инкремент на уровне процессора
		}()
	}

	wg.Wait()
	fmt.Println("Атомарный счётчик:", atomic.LoadInt64(&counter))
}
