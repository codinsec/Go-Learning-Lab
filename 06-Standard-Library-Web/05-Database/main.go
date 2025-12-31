package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/mattn/go-sqlite3" // SQLite driver (example)
)

// This program demonstrates database operations in Go

// User struct for database
type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	fmt.Println("=== Database Operations ===")
	fmt.Println()
	fmt.Println("Note: This example uses SQLite for demonstration.")
	fmt.Println("For PostgreSQL/MySQL, use appropriate drivers:")
	fmt.Println("  - PostgreSQL: github.com/lib/pq")
	fmt.Println("  - MySQL: github.com/go-sql-driver/mysql")
	fmt.Println()

	// 1. Database connection
	fmt.Println("1. Database Connection:")
	db, err := sql.Open("sqlite3", ":memory:") // In-memory database
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("   Database connected successfully")
	fmt.Println()

	// 2. Create table
	fmt.Println("2. Create Table:")
	createTable(db)
	fmt.Println()

	// 3. Insert data
	fmt.Println("3. Insert Data:")
	insertUser(db, "Alice", "alice@example.com")
	insertUser(db, "Bob", "bob@example.com")
	fmt.Println()

	// 4. Query single row
	fmt.Println("4. Query Single Row:")
	user := queryUser(db, 1)
	if user != nil {
		fmt.Printf("   User: %+v\n", *user)
	}
	fmt.Println()

	// 5. Query multiple rows
	fmt.Println("5. Query Multiple Rows:")
	users := queryAllUsers(db)
	for _, u := range users {
		fmt.Printf("   User: %+v\n", u)
	}
	fmt.Println()

	// 6. Prepared statements
	fmt.Println("6. Prepared Statements:")
	insertUserPrepared(db, "Charlie", "charlie@example.com")
	fmt.Println()

	// 7. Transactions
	fmt.Println("7. Transactions:")
	transactionExample(db)
	fmt.Println()

	// 8. Connection pooling
	fmt.Println("8. Connection Pooling:")
	fmt.Printf("   SetMaxOpenConns: %d\n", db.Stats().MaxOpenConnections)
	fmt.Println("   Connection pool is managed automatically")
	fmt.Println()

	// 9. Error handling
	fmt.Println("9. Error Handling:")
	handleDatabaseErrors(db)
	fmt.Println()

	// 10. Best practices
	fmt.Println("10. Best Practices:")
	fmt.Println("   - Always use prepared statements")
	fmt.Println("   - Use transactions for multiple operations")
	fmt.Println("   - Close rows after iteration")
	fmt.Println("   - Handle errors properly")
	fmt.Println("   - Use connection pooling")
	fmt.Println("   - Don't store passwords in code")
}

func createTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE
	)`
	
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("   Table 'users' created")
}

func insertUser(db *sql.DB, name, email string) {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	result, err := db.Exec(query, name, email)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	
	id, _ := result.LastInsertId()
	fmt.Printf("   User inserted with ID: %d\n", id)
}

func queryUser(db *sql.DB, id int) *User {
	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := db.QueryRow(query, id)
	
	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("   User not found")
		} else {
			fmt.Printf("   Error: %v\n", err)
		}
		return nil
	}
	
	return &user
}

func queryAllUsers(db *sql.DB) []User {
	query := "SELECT id, name, email FROM users"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	
	return users
}

func insertUserPrepared(db *sql.DB, name, email string) {
	stmt, err := db.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	
	result, err := stmt.Exec(name, email)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	
	id, _ := result.LastInsertId()
	fmt.Printf("   User inserted with prepared statement, ID: %d\n", id)
}

func transactionExample(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	
	// Rollback on error
	defer func() {
		if err != nil {
			tx.Rollback()
			fmt.Println("   Transaction rolled back")
			return
		}
		err = tx.Commit()
		if err != nil {
			fmt.Printf("   Commit error: %v\n", err)
		} else {
			fmt.Println("   Transaction committed")
		}
	}()
	
	// Multiple operations
	_, err = tx.Exec("INSERT INTO users (name, email) VALUES (?, ?)", "David", "david@example.com")
	if err != nil {
		return
	}
	
	_, err = tx.Exec("UPDATE users SET email = ? WHERE name = ?", "david.new@example.com", "David")
	if err != nil {
		return
	}
}

func handleDatabaseErrors(db *sql.DB) {
	// Query non-existent user
	user := queryUser(db, 999)
	if user == nil {
		fmt.Println("   Handled error: User not found")
	}
	
	// Try to insert duplicate email
	_, err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", "Duplicate", "alice@example.com")
	if err != nil {
		fmt.Printf("   Handled error: %v\n", err)
	}
}

