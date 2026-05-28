/*
Лабораторная работа №12: Агрегатные функции и группировка в SQL через Go
*/
package main

import (
	"database/sql"
	"fmt"
	"log"
)
import _ "github.com/lib/pq"

// Функция для наполнения базы данных
func seedDatabase(db *sql.DB) {
	// Создаем таблицу, если её нет
	createTable := `
	CREATE TABLE IF NOT EXISTS countries (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) UNIQUE NOT NULL,
		capital VARCHAR(100) NOT NULL,
		area INT NOT NULL,
		population BIGINT NOT NULL,
		continent VARCHAR(50) NOT NULL
	);`
	_, _ = db.Exec(createTable)

	// Cписком стран для проверки всех заданий
	countries := []struct {
		name, capital string
		area          int
		population    int64
		continent     string
	}{
		{"Россия", "Москва", 17125191, 146000000, "Европа"},
		{"Германия", "Берлин", 357022, 83000000, "Европа"},
		{"Франция", "Париж", 643801, 67000000, "Европа"},
		{"Ватикан", "Ватикан", 1, 800, "Европа"}, // Для минимальной площади
		{"Китай", "Пекин", 9596961, 1411000000, "Азия"},
		{"Индия", "Нью-Дели", 3287263, 1393000000, "Азия"},
		{"Пакистан", "Исламабад", 881912, 225000000, "Азия"}, // заканчивается на "стан"
		{"Азербайджан", "Баку", 86600, 10000000, "Азия"},     // заканчивается на "ан", но не "стан"
		{"Ливан", "Бейрут", 10452, 6000000, "Азия"},          // заканчивается на "ан", но не "стан"
		{"Бразилия", "Бразилиа", 8515767, 214000000, "Южная Америка"},
		{"Аргентина", "Буэнос-Айрес", 2780400, 45000000, "Южная Америка"},
		{"США", "Вашингтон", 9833517, 331000000, "Северная Америка"},
		{"Канада", "Оттава", 9984670, 38000000, "Северная Америка"},
		{"Нигерия", "Абуджа", 923768, 211000000, "Африка"}, // Плотность > 30, население > 1 млн
		{"Египет", "Каир", 1002450, 104000000, "Африка"},   // Плотность > 30, население > 1 млн
	}

	query := `INSERT INTO countries (name, capital, area, population, continent) 
	          VALUES ($1, $2, $3, $4, $5) ON CONFLICT (name) DO NOTHING;`
	for _, c := range countries {
		_, _ = db.Exec(query, c.name, c.capital, c.area, c.population, c.continent)
	}
}

// 1. Минимальная площадь стран
func task1(db *sql.DB) {
	var minArea int
	err := db.QueryRow("SELECT MIN(area) FROM countries").Scan(&minArea)
	if err != nil {
		log.Println("Ошибка задания 1:", err)
		return
	}
	fmt.Printf("Задание 1. Минимальная площадь среди всех стран: %d кв. км.\n", minArea)
}

// 2. Наибольшая по населению страна в Северной и Южной Америке
func task2(db *sql.DB) {
	var name string
	var population int64
	query := `SELECT name, population FROM countries 
	          WHERE continent IN ('Северная Америка', 'Южная Америка') 
	          ORDER BY population DESC LIMIT 1`
	err := db.QueryRow(query).Scan(&name, &population)
	if err != nil {
		log.Println("Ошибка задания 2:", err)
		return
	}
	fmt.Printf("Задание 2. Наибольшая по населению страна в Америках: %s (%d чел.)\n", name, population)
}

// 3. Среднее население стран
func task3(db *sql.DB) {
	var avgPopulation float64
	err := db.QueryRow("SELECT ROUND(AVG(population), 1) FROM countries").Scan(&avgPopulation)
	if err != nil {
		log.Println("Ошибка задания 3:", err)
		return
	}
	fmt.Printf("Задание 3. Среднее население стран: %.1f чел.\n", avgPopulation)
}

// 4. Количество стран на «ан», кроме «стан»
func task4(db *sql.DB) {
	var count int
	query := "SELECT COUNT(*) FROM countries WHERE name LIKE '%ан' AND name NOT LIKE '%стан'"
	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		log.Println("Ошибка задания 4:", err)
		return
	}
	fmt.Printf("Задание 4. Количество стран на 'ан' (кроме 'стан'): %d\n", count)
}

// 5. Количество континентов, где есть страны на букву «Р»
func task5(db *sql.DB) {
	var count int
	query := "SELECT COUNT(DISTINCT continent) FROM countries WHERE name LIKE 'Р%'"
	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		log.Println("Ошибка задания 5:", err)
		return
	}
	fmt.Printf("Задание 5. Количество континентов со странами на букву 'Р': %d\n", count)
}

// 6. Отношение площадей самой большой и маленькой стран
func task6(db *sql.DB) {
	var ratio float64
	// Используем NULLIF для защиты от деления на ноль в SQL
	query := "SELECT MAX(area)::FLOAT / NULLIF(MIN(area), 0) FROM countries"
	err := db.QueryRow(query).Scan(&ratio)
	if err != nil {
		log.Println("Ошибка задания 6:", err)
		return
	}
	fmt.Printf("Задание 6. Площадь самой большой страны больше самой маленькой в %.2f раз\n", ratio)
}

// 7. Количество стран с населением > 100 млн на каждом континенте
func task7(db *sql.DB) {
	query := `SELECT continent, COUNT(*) as cnt FROM countries 
	          WHERE population > 100000000 
	          GROUP BY continent 
	          ORDER BY cnt ASC`
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Ошибка задания 7:", err)
		return
	}
	defer rows.Close()

	fmt.Println("Задание 7. Количество стран с населением > 100 млн по континентам:")
	for rows.Next() {
		var continent string
		var count int
		_ = rows.Scan(&continent, &count)
		fmt.Printf(" - %s: %d стран(ы)\n", continent, count)
	}
}

// 8. Количество стран по количеству букв в названии
func task8(db *sql.DB) {
	query := `SELECT LENGTH(name) as len, COUNT(*) as cnt FROM countries 
	          GROUP BY LENGTH(name) 
	          ORDER BY cnt DESC`
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Ошибка задания 8:", err)
		return
	}
	defer rows.Close()

	fmt.Println("Задание 8. Группировка стран по длине названия:")
	for rows.Next() {
		var length, count int
		_ = rows.Scan(&length, &count)
		fmt.Printf(" - Длина названиий %d букв: %d стран(ы)\n", length, count)
	}
}

// 9. Прогноз населения через 20 лет (+10%)
func task9(db *sql.DB) {
	query := `SELECT continent, ROUND(SUM(population) * 1.1) FROM countries 
	          GROUP BY continent`
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Ошибка задания 9:", err)
		return
	}
	defer rows.Close()

	fmt.Println("Задание 9. Прогноз населения через 20 лет по континентам (+10%):")
	for rows.Next() {
		var continent string
		var predictedPopulation int64
		_ = rows.Scan(&continent, &predictedPopulation)
		fmt.Printf(" - %s: %d чел.\n", continent, predictedPopulation)
	}
}

// 10. Контененты, где разница площадей макс/мин не превышает 10 000 раз
func task10(db *sql.DB) {
	query := `SELECT continent FROM countries 
	          GROUP BY continent 
	          HAVING MAX(area)::FLOAT <= 10000.0 * MIN(area)::FLOAT`
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Ошибка задания 10:", err)
		return
	}
	defer rows.Close()

	fmt.Println("Задание 10. Континенты, где разница площадей стран не превышает 10 000 раз:")
	for rows.Next() {
		var continent string
		_ = rows.Scan(&continent)
		fmt.Printf(" - %s\n", continent)
	}
}

// 11. Средняя длина названия африканских стран
func task11(db *sql.DB) {
	var avgLength float64
	query := "SELECT AVG(LENGTH(name)) FROM countries WHERE continent = 'Африка'"
	err := db.QueryRow(query).Scan(&avgLength)
	if err != nil {
		log.Println("Ошибка задания 11:", err)
		return
	}
	fmt.Printf("Задание 11. Средняя длина названия стран Африки: %.2f букв\n", avgLength)
}

// 12. Континенты со средней плотностью населения > 30 среди крупных стран
func task12(db *sql.DB) {
	query := `SELECT continent, AVG(population::FLOAT / area) as density FROM countries 
	          WHERE population > 1000000 
	          GROUP BY continent 
	          HAVING AVG(population::FLOAT / area) > 30 
	          ORDER BY density DESC`
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Ошибка задания 12:", err)
		return
	}
	defer rows.Close()

	fmt.Println("Задание 12. Континенты со средней плотностью населения > 30 чел./кв.км:")
	for rows.Next() {
		var continent string
		var density float64
		_ = rows.Scan(&continent, &density)
		fmt.Printf(" - %s: средняя плотность %.2f чел./кв. км.\n", continent, density)
	}
}

func main() {
	// Подключение к локальной базе в Docker
	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("База данных недоступна: %v. Проверьте, запущен ли Docker-контейнер.", err)
	}

	// Заполнение БД
	seedDatabase(db)
	fmt.Println("База данных наполнена")
	fmt.Println("==================================================")

	// Вызов всех 12 функций по заданию
	task1(db)
	fmt.Println()
	task2(db)
	fmt.Println()
	task3(db)
	fmt.Println()
	task4(db)
	fmt.Println()
	task5(db)
	fmt.Println()
	task6(db)
	fmt.Println()
	task7(db)
	fmt.Println()
	task8(db)
	fmt.Println()
	task9(db)
	fmt.Println()
	task10(db)
	fmt.Println()
	task11(db)
	fmt.Println()
	task12(db)
}
