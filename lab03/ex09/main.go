/*
9. Вложенные структуры: Создайте структуру Student с полями Name, Group и
вложенной структурой Grades (массив оценок). Напишите метод для
вычисления среднего балла.
*/
package main

import "fmt"

type Grades struct {
	Marks [5]int
}

type Student struct {
	Name  string
	Group string
	Grades
}

func (s Student) AverageMark() float64 {
	sum := 0
	for _, mark := range s.Marks {
		sum += mark
	}
	return float64(sum) / float64(len(s.Marks))
}

func main() {
	g := Grades{Marks: [5]int{5, 4, 5, 3, 4}}
	student := Student{Name: "Алексей", Group: "ИВТ-21", Grades: g}

	fmt.Printf("Средний балл студента %s: %.2f\n", student.Name, student.AverageMark())
}
