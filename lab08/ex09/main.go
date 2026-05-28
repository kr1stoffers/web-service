/*
9. Пакет errors: Изучите функции errors.Is и errors.As. Напишите пример их
использования.
*/
package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("ресурс не найден")

type DatabaseError struct {
	Code int
}

func (e *DatabaseError) Error() string { return "ошибка базы данных" }

func getDatabaseRecord() error {
	// Имитируем оборачивание ошибки в цепочку (error wrapping)
	return fmt.Errorf("проблема при получении данных: %w", ErrNotFound)
}

func main() {
	err := getDatabaseRecord()

	// errors.Is проверяет, есть ли конкретная базовая ошибка в цепочке
	if errors.Is(err, ErrNotFound) {
		fmt.Println("errors.Is зафиксировал: Причина ошибки действительно ErrNotFound!")
	}

	// errors.As проверяет тип ошибки и распаковывает её в переменную
	dbErrOriginal := &DatabaseError{Code: 500}
	var targetErr *DatabaseError

	if errors.As(dbErrOriginal, &targetErr) {
		fmt.Println("errors.As зафиксировал совпадение типа! Код ошибки:", targetErr.Code)
	}
}
