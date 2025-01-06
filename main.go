package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func connectDB() (*sql.DB, error) {
	return sql.Open("postgres", "host=localhost port=5432 user=postgres password=pass dbname=go_demo sslmode=disable")
}

func insertUser(db *sql.DB, name, email string) {
	_, err := db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", name, email)
	if err != nil {
		log.Fatalf("Failed to insert user: %v", err)
	}
	fmt.Println("User inserted.")
}

func queryUsers(db *sql.DB) {
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		log.Fatalf("Failed to query users: %v", err)
	}
	defer rows.Close()

	fmt.Println("Users:")
	for rows.Next() {
		var id int
		var name, email string
		rows.Scan(&id, &name, &email)
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", id, name, email)
	}
}

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	insertUser(db, "Alice", "alice@example.com")
	queryUsers(db)
}
