package main

import (
	"fmt"
	"log/slog"
	"os"
	"temp/database"
)

func main() {
	// Подключение к БД
	db, err := database.DBconnection()
	if err != nil {
		slog.Error("Failed to connect to DB", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer db.Close()

	fmt.Println("Successfully connection to DB!")

	// Создание таблицы
	if err := database.CreateTable(db); err != nil {
		slog.Error("Failed create Table: %v", slog.String("error", err.Error()))
	}

	var name string
	var contribution int
	var percent float64

	fmt.Print("Insert a name: ")
	fmt.Scan(&name)

	fmt.Print("Insert a contribution: ")
	fmt.Scan(&contribution)

	fmt.Print("Insert a percent: ")
	fmt.Scan(&percent)

	// Внедрение данных в поля таблицы
	err = database.InsertIntoTable(db, name, contribution, percent)
	if err != nil {
		slog.Error("Can't insert row: %v", slog.String("error", err.Error()))
	}

	// Запрос на выборку данных из БД
	printUsers(db)
}
