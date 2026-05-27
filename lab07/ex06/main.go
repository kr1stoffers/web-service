/*
6. Гонка данных: Напишите программу, в которой несколько горутин
инкрементируют общую переменную без мьютекса. Запустите с флагом -race
и убедитесь, что гонка обнаружена. Затем исправьте.
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex // Мьютекс для исправления гонки данных
	counter := 0

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// ДЛЯ ИСПРАВЛЕНИЯ ГОНКИ: Раскомментируйте mu.Lock() и mu.Unlock()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Финальное значение счётчика:", counter)
}
