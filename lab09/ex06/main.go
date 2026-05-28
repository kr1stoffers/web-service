/*
6. Запись структуры в файл: Определите структуру Person с полями Name, Age.
Создайте несколько экземпляров и запишите их в файл в формате JSON
(используйте encoding/json). Затем прочитайте и выведите.
*/
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	people := []Person{
		{Name: "Дмитрий", Age: 20},
		{Name: "Елена", Age: 22},
	}

	// Запись в файл
	file, _ := os.Create("people.json")
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	_ = encoder.Encode(people)
	file.Close()

	// Чтение из файла
	readFile, _ := os.Open("people.json")
	defer readFile.Close()

	var decodedPeople []Person
	decoder := json.NewDecoder(readFile)
	_ = decoder.Decode(&decodedPeople)

	fmt.Println("Данные из JSON файла:")
	for _, p := range decodedPeople {
		fmt.Printf("Имя: %s, Возраст: %d\n", p.Name, p.Age)
	}
}
