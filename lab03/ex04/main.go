/*
4.  Библиотека: Создайте структуру Library, содержащую поле Books (срез
книг). Напишите метод AddBook(book Book) и метод PrintAll().
*/
package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Year   int
}

type Library struct {
	Books []Book
}

func (l *Library) AddBook(book Book) {
	l.Books = append(l.Books, book)
}

func (l Library) PrintAll() {
	for _, book := range l.Books {
		fmt.Printf("«%s» — %s (%d)\n", book.Title, book.Author, book.Year)
	}
}

func main() {
	lib := Library{}
	lib.AddBook(Book{"Мертвые души", "Н.В. Гоголь", 1842})
	lib.AddBook(Book{"Идиот", "Ф.М. Достоевский", 1869})
	lib.PrintAll()
}
