/*
9. Множественные интерфейсы: Создайте тип Robot, который реализует
интерфейсы Speaker и Walker. Продемонстрируйте использование.
*/
package main

import "fmt"

type Speaker interface {
	Speak()
}

type Walker interface {
	Walk()
}

type Robot struct {
	Model string
}

func (r Robot) Speak() {
	fmt.Printf("Робот %s говорит: Приветствую!\n", r.Model)
}
func (r Robot) Walk() { fmt.Printf("Робот %s шагает вперед.\n", r.Model) }

func main() {
	bot := Robot{Model: "xiaomi"}

	var sp Speaker = bot
	sp.Speak()

	var wk Walker = bot
	wk.Walk()
}
