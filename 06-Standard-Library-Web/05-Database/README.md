# Database

Learn about database operations in Go using the `database/sql` package.

## Learning Objectives

- Understand how to connect to databases
- Learn about executing SQL queries
- Master prepared statements
- Understand transactions
- Learn about connection pooling
- Understand error handling

## Concepts Covered

### Database Connection

```go
import (
    "database/sql"
    _ "github.com/lib/pq" // PostgreSQL driver
)

db, err := sql.Open("postgres", "connection string")
if err != nil {
    log.Fatal(err)
}
defer db.Close()

// Test connection
err = db.Ping()
```

### Executing Queries

**Exec (INSERT, UPDATE, DELETE):**
```go
result, err := db.Exec("INSERT INTO users (name) VALUES ($1)", "Alice")
id, _ := result.LastInsertId()
rowsAffected, _ := result.RowsAffected()
```

**QueryRow (single row):**
```go
var user User
err := db.QueryRow("SELECT name FROM users WHERE id = $1", id).Scan(&user.Name)
```

**Query (multiple rows):**
```go
rows, err := db.Query("SELECT id, name FROM users")
defer rows.Close()

for rows.Next() {
    var user User
    rows.Scan(&user.ID, &user.Name)
}
```

### Prepared Statements

```go
stmt, err := db.Prepare("INSERT INTO users (name) VALUES ($1)")
defer stmt.Close()

stmt.Exec("Alice")
stmt.Exec("Bob")
```

### Transactions

```go
tx, err := db.Begin()
defer func() {
    if err != nil {
        tx.Rollback()
    } else {
        tx.Commit()
    }
}()

tx.Exec("INSERT INTO users (name) VALUES ($1)", "Alice")
tx.Exec("UPDATE users SET name = $1 WHERE id = $2", "Bob", 1)
```

### Connection Pooling

```go
db.SetMaxOpenConns(25)
db.SetMaxIdleConns(5)
db.SetConnMaxLifetime(5 * time.Minute)
```

## Running the Example

```bash
# Navigate to this directory
cd 05-Database

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/06-standard-library-web/database

# Download dependencies
go mod tidy

# Run the program
go run main.go
```

## Running Tests

```bash
# Run all tests
go test

# Run tests with verbose output
go test -v

# Run tests with coverage
go test -cover
```

## Key Takeaways

1. **database/sql** - Standard interface for SQL databases
2. **Drivers** - Import driver with blank identifier
3. **Always close** - Close db, rows, statements
4. **Prepared statements** - More secure and efficient
5. **Transactions** - For multiple related operations
6. **Connection pooling** - Managed automatically

## Common Patterns

**Query single row:**
```go
var value string
err := db.QueryRow("SELECT name FROM users WHERE id = $1", id).Scan(&value)
```

**Query multiple rows:**
```go
rows, _ := db.Query("SELECT * FROM users")
defer rows.Close()
for rows.Next() {
    // Process row
}
```

**Transaction:**
```go
tx, _ := db.Begin()
defer tx.Rollback()
// Operations
tx.Commit()
```

## Best Practices

- **Use prepared statements** - Prevent SQL injection
- **Always close resources** - db, rows, statements
- **Handle errors** - Check all errors
- **Use transactions** - For related operations
- **Set connection limits** - Prevent resource exhaustion
- **Never store credentials** - Use environment variables

## Database Drivers

- **PostgreSQL**: `github.com/lib/pq`
- **MySQL**: `github.com/go-sql-driver/mysql`
- **SQLite**: `github.com/mattn/go-sqlite3`
- **SQL Server**: `github.com/denisenkom/go-mssqldb`

## Important Notes

- **Import drivers with `_`** - Blank identifier import
- **Placeholders vary** - PostgreSQL ($1), MySQL (?), SQLite (?)
- **Connection pooling** - Automatic, but can be configured
- **Context support** - Use context for cancellation

## Next Steps

After understanding database operations, proceed to:
- [06-Template-Engine](../06-Template-Engine/README.md) - Learn about template engines

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

