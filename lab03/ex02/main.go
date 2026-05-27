/*
2. Структура "Книга": Определите структуру Book с полями Title, Author, Year.
Создайте несколько экземпляров и выведите информацию.
*/
package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Year   int
}

func main() {
	b1 := Book{Title: "Капитанская дочка", Author: "А.С. Пушкин", Year: 1836}
	b2 := Book{Title: "1984", Author: "Дж. Оруэлл", Year: 1949}

	fmt.Println(b1)
	fmt.Println(b2)
}
