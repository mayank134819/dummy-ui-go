package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./users.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Query the users table
    rows, err := db.Query("SELECT id, name, email, password FROM users")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // Iterate over the rows and print the data
    for rows.Next() {
        var id int
        var name, email string
		var password string
        if err := rows.Scan(&id, &name, &email, &password); err != nil {
            log.Fatal(err)
        }
        fmt.Printf("ID: %d, Password: %s, Name: %s, Email: %s\n", id, password, name, email)
    }

    // Check for any errors from iterating over rows
    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
}
