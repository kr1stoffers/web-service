/*
7. Ошибки в горутинах: Запустите горутину, которая может вернуть ошибку. Как
правильно обработать ошибку из горутины? (Используйте канал для передачи ошибки.)
*/
package main

import (
	"errors"
	"fmt"
)

func asyncWorker(errCh chan<- error) {
	isFailed := true
	if isFailed {
		errCh <- errors.New("ошибка сети внутри горутины")
		return
	}
	errCh <- nil
}

func main() {
	errCh := make(chan error)
	go asyncWorker(errCh)

	err := <-errCh
	if err != nil {
		fmt.Println("Поймали ошибку из горутины в main:", err)
	}
}
