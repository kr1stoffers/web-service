/*
Лабораторная работа №13: Аналитика и обработка данных на стороне Go
ЧАСТЬ 1: Базовая загрузка и функции поиска с возвратом error
*/
package main

import (
	"database/sql"
	"fmt"
	"log"
	"sort"
	"unicode/utf8"

	_ "github.com/lib/pq"
)

// 1. Структура Country: Соответствует столбцам таблицы. Поле Валюта поддерживает NULL (Задание 7).
type Country struct {
	ID         int
	Name       string
	Capital    string
	Area       int
	Population int64
	Continent  string
	Currency   sql.NullString // Задание 7: Слой для обработки nullable-полей в Go
}

// Вспомогательная функция для автоматической подготовки Docker СУБД к тестам
func setupDatabase(db *sql.DB) {
	_, _ = db.Exec(`CREATE TABLE IF NOT EXISTS countries (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) UNIQUE NOT NULL,
		capital VARCHAR(100) NOT NULL,
		area INT NOT NULL,
		population BIGINT NOT NULL,
		continent VARCHAR(50) NOT NULL
	);`)

	_, _ = db.Exec(`ALTER TABLE countries ADD COLUMN IF NOT EXISTS currency VARCHAR(50);`)

	countries := []struct {
		name, capital string
		area          int
		population    int64
		continent     string
		currency      interface{}
	}{
		{"Россия", "Москва", 17125191, 146000000, "Европа", "Рубль"},
		{"Германия", "Берлин", 357022, 83000000, "Европа", "Евро"},
		{"Франция", "Париж", 643801, 67000000, "Европа", "Евро"},
		{"Италия", "Рим", 301230, 60000000, "Европа", "Евро"},
		{"Ватикан", "Ватикан", 1, 800, "Европа", nil}, // NULL валюта
		{"Китай", "Пекин", 9596961, 1411000000, "Азия", "Юань"},
		{"Индия", "Нью-Дели", 3287263, 1393000000, "Азия", "Рупия"},
		{"Япония", "Токио", 377975, 126000000, "Азия", "Иена"},
		{"Бразилия", "Бразилиа", 8515767, 214000000, "Южная Америка", "Реал"},
		{"Аргентина", "Буэнос-Айрес", 2780400, 45000000, "Южная Америка", "Песо"},
		{"США", "Вашингтон", 9833517, 331000000, "Северная Америка", "Доллар"},
		{"Канада", "Оттава", 9984670, 38000000, "Северная Америка", "Доллар"},
		{"Нигерия", "Абуджа", 923768, 211000000, "Африка", "Найра"},
	}

	query := `INSERT INTO countries (name, capital, area, population, continent, currency) 
	          VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (name) DO NOTHING;`
	for _, c := range countries {
		_, _ = db.Exec(query, c.name, c.capital, c.area, c.population, c.continent, c.currency)
	}
}

// 1. Функция GetAllCountries(): Загружает сырой срез всех стран из БД в ОЗУ
func GetAllCountries(db *sql.DB) ([]Country, error) {
	query := `SELECT id, name, capital, area, population, continent, currency FROM countries`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Country
	for rows.Next() {
		var c Country
		err := rows.Scan(&c.ID, &c.Name, &c.Capital, &c.Area, &c.Population, &c.Continent, &c.Currency)
		if err != nil {
			return nil, err
		}
		list = append(list, c)
	}
	return list, nil
}

// 2. Выборка по континенту: Фильтр выполняется силами Go на основе загруженных данных
func GetCountriesByContinent(countries []Country, continent string) ([]Country, error) {
	var result []Country
	for _, c := range countries {
		if c.Continent == continent {
			result = append(result, c)
		}
	}
	return result, nil
}

// 3. Поиск страны по названию: Выполняется силами Go (возвращает указатель и error)
func FindCountryByName(countries []Country, name string) (*Country, error) {
	for _, c := range countries {
		if c.Name == name {
			return &c, nil
		}
	}
	return nil, nil // Ошибок драйвера нет, объект просто отсутствует
}

// 4. Фильтрация на стороне Go (население > 50 миллионов)
func task4(countries []Country) {
	fmt.Println("Задание 4. Фильтрация стран с населением > 50 млн:")
	for _, c := range countries {
		if c.Population > 50000000 {
			fmt.Printf(" - Страна: %-10s | Население: %d чел.\n", c.Name, c.Population)
		}
	}
}

// 5. Сортировка по населению: Вывод топ-10 самых населённых стран
func task5(countries []Country) {
	// Делаем копию среза, чтобы не разрушить исходную последовательность
	sorted := make([]Country, len(countries))
	copy(sorted, countries)

	// Быстрая сортировка на стороне Go по убыванию поля Population
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Population > sorted[j].Population
	})

	fmt.Println("Задание 5. Топ-10 самых населённых стран (сортировка в Go):")
	for i := 0; i < 10 && i < len(sorted); i++ {
		fmt.Printf("  %d. %-10s — %d чел.\n", i+1, sorted[i].Name, sorted[i].Population)
	}
}

// 6. Статистика по континентам (Агрегация в Go без GROUP BY в SQL)
func task6(countries []Country) {
	type continentStats struct {
		totalPopulation int64
		totalArea       int64
		countryCount    int
	}

	statsMap := make(map[string]*continentStats)

	// Ручная агрегация в карту (имитация GROUP BY)
	for _, c := range countries {
		if _, exists := statsMap[c.Continent]; !exists {
			statsMap[c.Continent] = &continentStats{}
		}
		s := statsMap[c.Continent]
		s.totalPopulation += c.Population
		s.totalArea += int64(c.Area)
		s.countryCount++
	}

	fmt.Println("Задание 6. Сводный аналитический отчет по континентам (расчет в Go):")
	fmt.Printf("%-20s | %-16s | %-18s | %-12s\n", "Континент", "Общее население", "Средняя площадь", "Кол-во стран")
	fmt.Println("-------------------------------------------------------------------------")
	for continent, s := range statsMap {
		avgArea := float64(s.totalArea) / float64(s.countryCount)
		fmt.Printf("%-20s | %-16d | %-18.1f | %-12d\n", continent, s.totalPopulation, avgArea, s.countryCount)
	}
}

// 7. Обработка и вывод Nullable-полей (Задание 7)
func task7(countries []Country) {
	fmt.Println("Задание 7. Проверка nullable-поля Валюта:")
	for _, c := range countries {
		currencyValue := "NULL (Данные отсутствуют)"
		if c.Currency.Valid {
			currencyValue = c.Currency.String
		}
		fmt.Printf(" - %-10s | Валюта: %s\n", c.Name, currencyValue)
	}
}

// 8. Самая длинная и самая короткая столица по количеству символов Unicode
func task8(countries []Country) {
	if len(countries) == 0 {
		return
	}

	maxCountry := countries[0]
	minCountry := countries[0]

	for _, c := range countries {
		// Обязательно utf8.RuneCountInString для корректной обработки русских букв
		if utf8.RuneCountInString(c.Capital) > utf8.RuneCountInString(maxCountry.Capital) {
			maxCountry = c
		}
		if utf8.RuneCountInString(c.Capital) < utf8.RuneCountInString(minCountry.Capital) {
			minCountry = c
		}
	}

	fmt.Println("Задание 8. Анализ длины названий столиц:")
	fmt.Printf(" - Самая длинная столица: %s (Страна: %s, Букв: %d)\n", maxCountry.Capital, maxCountry.Name, utf8.RuneCountInString(maxCountry.Capital))
	fmt.Printf(" - Самая короткая столица: %s (Страна: %s, Букв: %d)\n", minCountry.Capital, minCountry.Name, utf8.RuneCountInString(minCountry.Capital))
}

// 9. Группировка стран по первой букве названия (Ключ — rune)
func task9(countries []Country) {
	groups := make(map[rune][]Country)

	for _, c := range countries {
		runes := []rune(c.Name)
		if len(runes) > 0 {
			firstLetter := runes[0]
			groups[firstLetter] = append(groups[firstLetter], c)
		}
	}

	fmt.Println("Задание 9. Количество стран по первой букве названия:")
	for letter, list := range groups {
		fmt.Printf(" - Буква '%c': %d страна(ы)\n", letter, len(list))
	}
}

// 10. Континенты с максимальной плотностью населения
func task10(countries []Country) {
	maxDensityMap := make(map[string]Country)

	densityCalc := func(c Country) float64 {
		if c.Area == 0 {
			return 0
		}
		return float64(c.Population) / float64(c.Area)
	}

	for _, c := range countries {
		currentDensity := densityCalc(c)
		bestCountry, exists := maxDensityMap[c.Continent]
		if !exists || currentDensity > densityCalc(bestCountry) {
			maxDensityMap[c.Continent] = c
		}
	}

	fmt.Println("Задание 10. Страны с максимальной плотностью населения на каждом континенте:")
	for continent, c := range maxDensityMap {
		fmt.Printf(" - %-20s: %-10s (Макс. Плотность: %.2f чел./кв. км.)\n", continent, c.Name, densityCalc(c))
	}
}

func main() {
	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка конфигурации: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Ошибка сети Docker: %v. Убедитесь, что контейнер запущен.", err)
	}

	// Готовим сырые данные в базе данных
	setupDatabase(db)
	fmt.Println("База данных успешно подготовлена.")
	fmt.Println("==================================================\n")

	// Вызов Задания 1: Читаем всю таблицу в ОЗУ один раз
	allCountries, err := GetAllCountries(db)
	if err != nil {
		log.Fatalf("Ошибка загрузки данных: %v", err)
	}
	fmt.Printf("Задание 1. Срез всех стран загружен из СУБД в ОЗУ. Всего элементов: %d\n\n", len(allCountries))

	// Вызов Задания 2: Выборка по континенту (Европа) на стороне Go
	europe, err := GetCountriesByContinent(allCountries, "Европа")
	if err == nil {
		fmt.Printf("Задание 2. Из загруженных данных отфильтрованы страны Европы. Найдено: %d\n\n", len(europe))
	}

	// Вызов Задания 3: Поиск по названию
	fmt.Println("Задание 3. Тестирование функций поиска в RAM:")
	found, _ := FindCountryByName(allCountries, "Россия")
	if found != nil {
		fmt.Printf(" - Поиск 'Россия': Найдено. Столица: %s\n", found.Capital)
	}
	notFound, _ := FindCountryByName(allCountries, "что-то")
	if notFound == nil {
		fmt.Println(" - Поиск 'что-то': вернул nil (страна отсутствует).")
	}
	fmt.Println()

	// Вызов остальных аналитических функций (Задания 4–10)
	task4(allCountries)
	fmt.Println()

	task5(allCountries)
	fmt.Println()

	task6(allCountries)
	fmt.Println()

	task7(allCountries)
	fmt.Println()

	task8(allCountries)
	fmt.Println()

	task9(allCountries)
	fmt.Println()

	task10(allCountries)
}
