/*
Лабораторная работа №11: Взаимодействие с СУБД PostgreSQL в Go
задания 2, 3, 4, 5, 6.
*/
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

import _ "github.com/lib/pq"

// Структура для представления страны (Задание 5)
type Country struct {
	ID         int
	Name       string
	Capital    string
	Area       int
	Population int64
	Continent  string
}

// Задание 4: Функция добавления новой страны
func insertCountry(db *sql.DB, name, capital string, area int, population int64, continent string) error {
	query := `
		INSERT INTO countries (name, capital, area, population, continent)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (name) DO NOTHING;` // Защита от дубликатов при повторном запуске

	_, err := db.Exec(query, name, capital, area, population, continent)
	return err
}

// Задание 5: Функция выборки стран по континенту
func getCountriesByContinent(db *sql.DB, continent string) ([]Country, error) {
	query := `SELECT id, name, capital, area, population, continent FROM countries WHERE continent = $1`
	rows, err := db.Query(query, continent)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var countries []Country
	for rows.Next() {
		var c Country
		err := rows.Scan(&c.ID, &c.Name, &c.Capital, &c.Area, &c.Population, &c.Continent)
		if err != nil {
			return nil, err
		}
		countries = append(countries, c)
	}

	// Задание 6: Обработка ситуации, когда строк в таблице нет
	if len(countries) == 0 {
		return nil, nil // Возвращаем пустой срез без ошибки
	}

	return countries, nil
}

// Задание 3: Динамическое создание таблицы с проверкой существования
func ensureTableExists(db *sql.DB) error {
	var exists bool
	// Проверяем наличие таблицы в системном каталоге PostgreSQL
	checkQuery := `SELECT EXISTS (
		SELECT FROM information_schema.tables 
		WHERE table_name = 'countries'
	);`

	err := db.QueryRow(checkQuery).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		fmt.Println("Таблица уже существует.")
		return nil
	}

	fmt.Println("Таблицы нет. Создаем таблицу countries...")
	createQuery := `
	CREATE TABLE countries (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) UNIQUE NOT NULL,
		capital VARCHAR(100) NOT NULL,
		area INT NOT NULL,
		population BIGINT NOT NULL,
		continent VARCHAR(50) NOT NULL
	);`

	_, err = db.Exec(createQuery)
	return err
}

func main() {
	// Задание 2: Модификация строки подключения (считывание из переменной окружения)
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		// Значение по умолчанию для локального Docker-контейнера
		connStr = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
		fmt.Println("Переменная DATABASE_URL не задана. Используется значение по умолчанию.")
	} else {
		fmt.Println("Строка подключения успешно считана из DATABASE_URL.")
	}

	// Пул соединений с БД
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка конфигурации БД: %v", err)
	}
	defer db.Close()

	// Проверка подключения к базе данных
	err = db.Ping()
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	fmt.Println("Успешно подключились к PostgreSQL!")

	// Вызов Задания 3: Проверка и создание таблицы
	err = ensureTableExists(db)
	if err != nil {
		log.Fatalf("Ошибка при подготовке таблицы: %v", err)
	}

	// Вызов Задания 4: Добавление Бразилии
	err = insertCountry(db, "Бразилия", "Бразилиа", 8515767, 214300000, "Южная Америка")
	if err != nil {
		log.Printf("Ошибка добавления Бразилии: %v", err)
	} else {
		fmt.Println("Страна 'Бразилия' успешно обработана.")
	}

	// Тестовая европейская страна для проверки Задания 5
	_ = insertCountry(db, "Франция", "Париж", 643801, 67500000, "Европа")

	// Вызов Задания 5 и 6: Выборка стран по континенту "Европа"
	fmt.Println("\n--- Выборка стран по континенту 'Европа' ---")
	europeCountries, err := getCountriesByContinent(db, "Европа")
	if err != nil {
		log.Fatalf("Ошибка выполнения запроса: %v", err)
	}

	if europeCountries == nil {
		// Реализация Задания 6: Понятное сообщение вместо ошибки
		fmt.Println("В таблице нет данных по этому континенту.")
	} else {
		for _, c := range europeCountries {
			fmt.Printf("ID: %d | %s (Столица: %s) | Население: %d | Континент: %s\n",
				c.ID, c.Name, c.Capital, c.Population, c.Continent)
		}
	}

	// Проверка Задания 6: Выборка по континенту, которого нет в БД
	fmt.Println("\n--- Выборка стран по континенту 'Антарктида' ---")
	antarcticaCountries, _ := getCountriesByContinent(db, "Антарктида")
	if antarcticaCountries == nil {
		fmt.Println("В таблице нет данных.")
	}
}
