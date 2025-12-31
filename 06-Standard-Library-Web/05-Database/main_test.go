package main

import (
	"database/sql"
	"testing"
	_ "github.com/mattn/go-sqlite3"
)

func TestDatabaseConnection(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}
}

func TestCreateTable(t *testing.T) {
	db, _ := sql.Open("sqlite3", ":memory:")
	defer db.Close()

	createTable(db)

	// Verify table exists by querying
	rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table' AND name='users'")
	if err != nil {
		t.Fatalf("Failed to query: %v", err)
	}
	defer rows.Close()

	if !rows.Next() {
		t.Error("Table 'users' should exist")
	}
}

func TestInsertAndQuery(t *testing.T) {
	db, _ := sql.Open("sqlite3", ":memory:")
	defer db.Close()

	createTable(db)
	insertUser(db, "Test User", "test@example.com")

	user := queryUser(db, 1)
	if user == nil {
		t.Fatal("User should exist")
	}

	if user.Name != "Test User" {
		t.Errorf("Expected name 'Test User', got %s", user.Name)
	}

	if user.Email != "test@example.com" {
		t.Errorf("Expected email 'test@example.com', got %s", user.Email)
	}
}

func TestQueryAllUsers(t *testing.T) {
	db, _ := sql.Open("sqlite3", ":memory:")
	defer db.Close()

	createTable(db)
	insertUser(db, "User1", "user1@example.com")
	insertUser(db, "User2", "user2@example.com")

	users := queryAllUsers(db)
	if len(users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(users))
	}
}

func TestPreparedStatement(t *testing.T) {
	db, _ := sql.Open("sqlite3", ":memory:")
	defer db.Close()

	createTable(db)
	insertUserPrepared(db, "Prepared User", "prepared@example.com")

	user := queryUser(db, 1)
	if user == nil {
		t.Fatal("User should exist")
	}

	if user.Name != "Prepared User" {
		t.Errorf("Expected name 'Prepared User', got %s", user.Name)
	}
}

func TestTransaction(t *testing.T) {
	db, _ := sql.Open("sqlite3", ":memory:")
	defer db.Close()

	createTable(db)
	
	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	}

	_, err = tx.Exec("INSERT INTO users (name, email) VALUES (?, ?)", "Tx User", "tx@example.com")
	if err != nil {
		tx.Rollback()
		t.Fatalf("Failed to insert: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		t.Fatalf("Failed to commit: %v", err)
	}

	user := queryUser(db, 1)
	if user == nil {
		t.Error("User should exist after transaction")
	}
}

