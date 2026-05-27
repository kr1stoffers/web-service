/*
10. Дополнительно: Изучите пакет reflect и напишите функцию, которая
выводит тип и значение переменной, переданной как interface{}.
*/
package main

import (
	"fmt"
	"reflect"
)

func printTypeAndValue(i interface{}) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	fmt.Printf("Тип: %s, Значение: %v\n", t, v)
}

func main() {
	printTypeAndValue(42)
	printTypeAndValue("вот")
	printTypeAndValue(3.14)
}
