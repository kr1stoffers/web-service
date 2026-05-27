/*
4. Сортировка: Используя интерфейс sort.Interface, реализуйте сортировку
среза структур по одному из полей (например, Person по возрасту).
*/
package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	people := []Person{
		{"Анна", 25},
		{"Иван", 19},
		{"Ольга", 32},
	}

	sort.Sort(ByAge(people))
	fmt.Println(people)
}
