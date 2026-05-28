/*
Лабораторная работа №14: Безопасность БД, Подготовленные запросы и Транзакции в Go
*/
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Задание 8: Добавлен столбец ID int в структуру Country
type Country struct {
	ID         int
	Name       string
	Capital    string
	Area       int
	Population int64
	Continent  string
}

// Вспомогательная функция настройки структуры таблицы
func initDatabase(db *sql.DB) {
	// Таблица с автоинкрементным SERIAL PRIMARY KEY и UNIQUE ограничением для имени
	query := `
	CREATE TABLE IF NOT EXISTS countries_lab14 (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) UNIQUE NOT NULL,
		capital VARCHAR(100) NOT NULL,
		area INT NOT NULL,
		population BIGINT NOT NULL,
		continent VARCHAR(50) NOT NULL
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Ошибка инициализации таблицы: %v", err)
	}

	// Очищаем таблицу перед каждым демонстрационным запуском
	_, _ = db.Exec("TRUNCATE TABLE countries_lab14 RESTART IDENTITY;")
}

// 1. Добавление страны с безопасными плейсхолдерами
func AddCountry(db *sql.DB, c Country) error {
	query := `INSERT INTO countries_lab14 (name, capital, area, population, continent) VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(query, c.Name, c.Capital, c.Area, c.Population, c.Continent)
	return err
}

// 2. Обновление столицы через плейсхолдеры
func UpdateCapital(db *sql.DB, countryName, newCapital string) error {
	query := `UPDATE countries_lab14 SET capital = $1 WHERE name = $2`
	res, err := db.Exec(query, newCapital, countryName)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("страна %s не найдена", countryName)
	}
	return nil
}

// 3. Удаление стран по континенту с возвратом количества удаленных строк
func DeleteCountriesByContinent(db *sql.DB, continent string) (int64, error) {
	query := `DELETE FROM countries_lab14 WHERE continent = $1`
	res, err := db.Exec(query, continent)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// 4. Поиск по части названия (Безопасный LIKE через конкатенацию параметра)
func SearchCountriesByName(db *sql.DB, pattern string) ([]Country, error) {
	query := `SELECT id, name, capital, area, population, continent FROM countries_lab14 WHERE name LIKE $1`
	// Безопасное формирование паттерна: передаем знаки % внутри самого аргумента
	rows, err := db.Query(query, "%"+pattern+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Country
	for rows.Next() {
		var c Country
		err := rows.Scan(&c.ID, &c.Name, &c.Capital, &c.Area, &c.Population, &c.Continent)
		if err != nil {
			return nil, err
		}
		result = append(result, c)
	}
	return result, nil
}

// 5. Транзакция: обновление населения нескольких стран (атомарный откат при отсутствии любой страны)
func UpdatePopulationsTx(db *sql.DB, updates map[string]int64) error {
	tx, err := db.Begin() // Шаг 1. Открываем транзакцию
	if err != nil {
		return err
	}
	// В случае паники или непредвиденного выхода гарантируем откат
	defer tx.Rollback()

	query := `UPDATE countries_lab14 SET population = $1 WHERE name = $2`
	for countryName, newPop := range updates {
		res, err := tx.Exec(query, newPop, countryName)
		if err != nil {
			return err // Ошибка SQL приведет к автоматическому Rollback через defer
		}
		rowsAffected, _ := res.RowsAffected()
		if rowsAffected == 0 {
			// Если страна не найдена — принудительно возвращаем ошибку для отмены всей пачки
			return fmt.Errorf("страна '%s' не найдена, транзакция полностью отменяется", countryName)
		}
	}

	return tx.Commit() // Шаг 2. Если всё прошло успешно — фиксируем изменения в БД
}

// 6. Подготовленный запрос для массовой оптимизированной вставки через Prepare
func MassInsertPrepare(db *sql.DB, list []Country) error {
	// Компилируем и подготавливаем шаблон запроса на стороне СУБД один раз
	stmt, err := db.Prepare(`INSERT INTO countries_lab14 (name, capital, area, population, continent) VALUES ($1, $2, $3, $4, $5)`)
	if err != nil {
		return err
	}
	defer stmt.Close() // Обязательно освобождаем ресурс шаблона

	// Выполняем вставку в цикле, отправляя только сырые параметры
	for _, c := range list {
		_, err := stmt.Exec(c.Name, c.Capital, c.Area, c.Population, c.Continent)
		if err != nil {
			return err
		}
	}
	return nil
}

// 7.1 Уязвимая функция (Конкатенация строк) — ПРИМЕР КАК ДЕЛАТЬ НЕЛЬЗЯ
func UnsafeSearchCountries(db *sql.DB, injectionPattern string) ([]Country, error) {
	// ДЫРА В БЕЗОПАСНОСТИ: строка склеивается напрямую через плюсы
	query := "SELECT id, name, capital, area, population, continent FROM countries_lab14 WHERE name = '" + injectionPattern + "'"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Country
	for rows.Next() {
		var c Country
		_ = rows.Scan(&c.ID, &c.Name, &c.Capital, &c.Area, &c.Population, &c.Continent)
		result = append(result, c)
	}
	return result, nil
}

// 8. Получение последнего вставленного ID через RETURNING id (PostgreSQL специфика)
func AddCountryReturnID(db *sql.DB, c Country) (int, error) {
	query := `INSERT INTO countries_lab14 (name, capital, area, population, continent) 
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var lastInsertedID int
	// QueryRow идеален, так как RETURNING возвращает ровно одно значение в одну строку
	err := db.QueryRow(query, c.Name, c.Capital, c.Area, c.Population, c.Continent).Scan(&lastInsertedID)
	if err != nil {
		return 0, err
	}
	return lastInsertedID, nil
}

// 9. Демонстрация транзакции с ошибкой дублирования
func DemonstrateTxError(db *sql.DB, useTransaction bool) {
	c1 := Country{Name: "Испания", Capital: "Мадрид", Area: 505990, Population: 47000000, Continent: "Европа"}
	c2 := Country{Name: "Испания", Capital: "Дубликат", Area: 100, Population: 100, Continent: "Европа"} // Вызовет ошибку уникальности

	if useTransaction {
		tx, _ := db.Begin()
		defer tx.Rollback()

		_, err1 := tx.Exec(`INSERT INTO countries_lab14 (name, capital, area, population, continent) VALUES ($1, $2, $3, $4, $5)`, c1.Name, c1.Capital, c1.Area, c1.Population, c1.Continent)
		_, err2 := tx.Exec(`INSERT INTO countries_lab14 (name, capital, area, population, continent) VALUES ($1, $2, $3, $4, $5)`, c2.Name, c2.Capital, c2.Area, c2.Population, c2.Continent)

		if err1 != nil || err2 != nil {
			fmt.Println("  [Транзакция]: Обнаружена ошибка дубликата! Откатываем всё назад.")
			return // Происходит автоматический Rollback
		}
		_ = tx.Commit()
	} else {
		// Без транзакции (Обычные независимые Exec)
		_, _ = db.Exec(`INSERT INTO countries_lab14 (name, capital, area, population, continent) VALUES ($1, $2, $3, $4, $5)`, c1.Name, c1.Capital, c1.Area, c1.Population, c1.Continent)
		_, err2 := db.Exec(`INSERT INTO countries_lab14 (name, capital, area, population, continent) VALUES ($1, $2, $3, $4, $5)`, c2.Name, c2.Capital, c2.Area, c2.Population, c2.Continent)
		if err2 != nil {
			fmt.Println("  [Без транзакции]: Произошла ошибка дублирования второй страны.")
		}
	}
}

// 10. Обновление с математическим условием через UPDATE параметры
func IncreasePopulation(db *sql.DB, minPopulation int64, percent float64) error {
	// Рассчитываем коэффициент прироста: например, если percent = 10.0, коэффициент = 1.10
	factor := 1.0 + (percent / 100.0)
	query := `UPDATE countries_lab14 SET population = ROUND(population * $1) WHERE population > $2`
	_, err := db.Exec(query, factor, minPopulation)
	return err
}

func main() {
	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка конфигурации: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("База данных в Docker недоступна: %v", err)
	}

	initDatabase(db)
	fmt.Println("База данных готова. Начинаем демонстрацию выполнения лабораторной работы.")
	fmt.Println("=========================================================================")

	// Демонстрация Задания 1: Вставка и проверка уникальности
	fmt.Println("Задание 1. Тестирование AddCountry:")
	cRussia := Country{Name: "Россия", Capital: "Москва", Area: 17125191, Population: 146000000, Continent: "Европа"}
	err = AddCountry(db, cRussia)
	if err == nil {
		fmt.Println("  - Страна 'Россия' успешно добавлена.")
	}
	// Пытаемся вставить дубликат
	errDuplicate := AddCountry(db, cRussia)
	if errDuplicate != nil {
		fmt.Printf("  - Обработана ошибка уникальности (ОЖИДАЕМО): %v\n", errDuplicate)
	}

	// Демонстрация Задания 2: Обновление столицы
	fmt.Println("\nЗадание 2. Тестирование UpdateCapital:")
	err = UpdateCapital(db, "Россия", "Санкт-Петербург")
	if err == nil {
		fmt.Println("  - Столица России успешно изменена на Санкт-Петербург.")
	}

	// Демонстрация Задания 6: Массовая вставка Prepare
	fmt.Println("\nЗадание 6. Массовая вставка 5 стран через Prepare:")
	batch := []Country{
		{0, "Китай", "Пекин", 9596961, 1411000000, "Азия"},
		{0, "Индия", "Нью-Дели", 3287263, 1393000000, "Азия"},
		{0, "США", "Вашингтон", 9833517, 331000000, "Северная Америка"},
		{0, "Бразилия", "Бразилиа", 8515767, 214000000, "Южная Америка"},
		{0, "Египет", "Каир", 1002450, 104000000, "Африка"},
	}
	err = MassInsertPrepare(db, batch)
	if err == nil {
		fmt.Println("  - Пакет из 5 стран успешно добавлен через скомпилированный Prepare statement.")
	}

	// Демонстрация Задания 4: Безопасный поиск по подстроке
	fmt.Println("\nЗадание 4. Безопасный поиск SearchCountriesByName (ищем подстроку 'ия'):")
	foundList, _ := SearchCountriesByName(db, "ия")
	for _, f := range foundList {
		fmt.Printf("  - Найдено: %s (Континент: %s)\n", f.Name, f.Continent)
	}

	// Демонстрация Задания 7: Взлом уязвимой функции через инъекцию
	fmt.Println("\nЗадание 7. Демонстрация хакерской SQL-инъекции:")
	// Хакер передает специальную строчку, ломающую кавычки и добавляющую условие OR TRUE
	hackerPayload := "ЛюбаяСтрока' OR '1'='1"
	injectedResult, _ := UnsafeSearchCountries(db, hackerPayload)
	fmt.Printf("  - Уязвимая функция из-за подмены выдала ВСЕ страны базы данных (Кол-во: %d) вместо одной!\n", len(injectedResult))
	fmt.Println("  - ИСПРАВЛЕНИЕ: Функция SearchCountriesByName из Задания 4 использует плейсхолдеры и полностью защищена.")

	// Демонстрация Задания 8: Получение сгенерированного ID
	fmt.Println("\nЗадание 8. Получение ID через RETURNING id:")
	cFrance := Country{Name: "Франция", Capital: "Париж", Area: 643801, Population: 67000000, Continent: "Европа"}
	generatedID, _ := AddCountryReturnID(db, cFrance)
	fmt.Printf("  - Страна 'Франция' успешно сохранена. База данных присвоила ей автоинкрементный ID: %d\n", generatedID)

	// Демонстрация Задания 5: Транзакция со сбоем (одна страна отсутствует)
	fmt.Println("\nЗадание 5. Тестирование транзакции UpdatePopulationsTx (одна страна отсутствует):")
	updates := map[string]int64{
		"Китай":     1500000000,
		"Атлантида": 999, // Этой страны нет в базе данных
	}
	errTx := UpdatePopulationsTx(db, updates)
	if errTx != nil {
		fmt.Printf("  - Транзакция отменена (ОЖИДАЕМО): %v\n", errTx)
	}

	// Демонстрация Задания 9: Сравнение поведения с транзакцией и без при ошибках дублирования
	fmt.Println("\nЗадание 9. Разница поведения транзакции при ошибках дублирования:")
	// 9.1 Без транзакции
	DemonstrateTxError(db, false)
	resUnsafe, _ := SearchCountriesByName(db, "Испания")
	fmt.Printf("  - [Без транзакции]: Первая страна 'Испания' сохранилась в БД (Кол-во: %d), несмотря на ошибку второй.\n", len(resUnsafe))

	// Перезапускаем чистую базу для честного сравнения транзакции
	initDatabase(db)
	_ = AddCountry(db, cRussia) // Возвращаем Россию для структуры данных

	// 9.2 С транзакцией
	DemonstrateTxError(db, true)
	resSafe, _ := SearchCountriesByName(db, "Испания")
	fmt.Printf("  - [С транзакцией]: Из-за атомарности первая страна НЕ сохранилась в БД (Кол-во: %d). База чиста!\n", len(resSafe))

	// Демонстрация Задания 10: Обновление населения с условием (+15% для стран > 100 млн)
	fmt.Println("\nЗадание 10. Обновление с условием IncreasePopulation (+15% для стран с населением > 100 млн):")
	_ = MassInsertPrepare(db, batch) // возвращаем пачку стран в чистую базу
	err = IncreasePopulation(db, 100000000, 15.0)
	if err == nil {
		fmt.Println("  - Население крупных стран успешно проиндексировано на 15%.")
	}

	// Демонстрация Задания 3: Удаление по континенту
	fmt.Println("\nЗадание 3. Тестирование DeleteCountriesByContinent для континента 'Азия':")
	deletedCount, _ := DeleteCountriesByContinent(db, "Азия")
	fmt.Printf("  - Операция выполнена. Из базы успешно удалено строк: %d\n", deletedCount)
}
