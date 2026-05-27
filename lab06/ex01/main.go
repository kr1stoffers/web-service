/*
1. Интерфейс Shape: Определите интерфейс Shape с методами Area() float64 и Perimeter() float64.
2. Список фигур: Создайте срез Shape и выведите площадь и периметр каждой фигуры.
3. Интерфейс Stringer: Реализуйте метод String() для своих структур.
*/
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }
func (c Circle) String() string     { return fmt.Sprintf("Круг (R=%.2f)", c.Radius) }

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }
func (r Rectangle) String() string {
	return fmt.Sprintf("Прямоугольник (%.2fx%.2f)", r.Width, r.Height)
}

func main() {
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 6},
	}

	for _, shape := range shapes {
		fmt.Println(shape)
		fmt.Printf("Площадь: %.2f, Периметр: %.2f\n\n", shape.Area(), shape.Perimeter())
	}
}
