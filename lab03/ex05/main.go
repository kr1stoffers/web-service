/*
5. Телефонная книга: Используя карту, реализуйте телефонную книгу: ключ –
имя (строка), значение – номер (строка). Напишите функции Add, Get, Delete.
*/
package main

import "fmt"

var phoneBook = make(map[string]string)

func Add(name, phone string) {
	phoneBook[name] = phone
}

func Get(name string) (string, bool) {
	phone, exists := phoneBook[name]
	return phone, exists
}

func Delete(name string) {
	delete(phoneBook, name)
}

func main() {
	Add("Иван", "+79991112233")
	Add("Анна", "+79994445566")

	if phone, ok := Get("Иван"); ok {
		fmt.Println("Номер Ивана:", phone)
	}

	Delete("Иван")

	_, ok := Get("Иван")
	fmt.Println("Иван существует?", ok)
}
