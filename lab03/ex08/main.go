/*
8. Сравнение структур: Определите, можно ли сравнивать две структуры одного
типа с помощью ==. Проверьте на примере.
*/
package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	p1 := Point{X: 5, Y: 10}
	p2 := Point{X: 5, Y: 10}
	p3 := Point{X: 1, Y: 2}

	fmt.Println("p1 == p2:", p1 == p2)
	fmt.Println("p1 == p3:", p1 == p3)
}
