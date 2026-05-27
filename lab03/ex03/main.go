/*
3. Метод для книги: Добавьте метод Age(), возвращающий возраст книги
(текущий год минус год издания). Используйте time.Now().Year().
*/
package main

import (
	"fmt"
	"time"
)

type Book struct {
	Title  string
	Author string
	Year   int
}

func (b Book) Age() int {
	return time.Now().Year() - b.Year
}

func main() {
	b := Book{Title: "Преступление и наказание", Author: "Ф.М. Достоевский", Year: 1866}
	fmt.Println(b.Age())
}
