package database

import (
	"database/sql"
	"fmt"
	"log/slog"
)

func DBconnection() (*sql.DB, error) {
	connStr := "host=localhost port=8080 user=postgres password=DBPassword dbname=DBName sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging DB: %w", err)
	}

	return db, nil
}

func CreateTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users(
	  	   		   id SERIAL PRIMARY KEY,
	 			 name VARCHAR(50) NOT NULL,
		 contribution INT NOT NULL,
		 	  percent DECIMAL NOT NULL
	);
	`)
	return err
}

func InsertIntoTable(db *sql.DB, name string, contribution int, percent float64) error {
	query := `
		INSERT INTO users (name, contribution, percent)
		VALUES ($1, $2, $3)
	`
	_, err := db.Exec(query, name, contribution, percent)
	if err != nil {
		return fmt.Errorf("failed to insert into table: %w", err)
	}
	return nil
}

func printUsers(db *sql.DB) {
	rows, err := db.Query(`SELECT * FROM users`)
	if err != nil {
		slog.Error("Failed to query users", slog.String("error", err.Error()))
		return
	}
	defer rows.Close()

	fmt.Println("Users table: ")
	for rows.Next() {
		
	}

}
