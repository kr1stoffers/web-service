/*
10. Дополнительно: Реализуйте функцию compose, которая принимает две
функции f и g и возвращает функцию h(x) = f(g(x)).
*/
package main

import "fmt"

func compose(f func(int) int, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

func main() {
	addTwo := func(x int) int { return x + 2 }
	multiplyThree := func(x int) int { return x * 3 }

	h := compose(addTwo, multiplyThree)
	fmt.Println(h(5))
}
