package database

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

var db *sql.DB

// InitializeDB initializes the SQLite database and creates the user table if not exists.
func InitializeDB() {
	var err error
	db, err = sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal("Failed to open database: ", err)
	}


	// Create the users table if it doesn't exist
	createTableQuery := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Failed to create table: ", err)
	}

	log.Println("Database initialized successfully and users table is ready.")
}

// CreateUser inserts a new user into the database
func CreateUser(name, email, password string) error {
	_, err := db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", name, email, password)
	log.Println("Any error in creating data in database.", err)
	return err
}

// GetUser fetches a user by email and password
func GetUser(email, password string) (*User, error) {
	row := db.QueryRow("SELECT id, name, email FROM users WHERE email = ? AND password = ?", email, password)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// User represents a user in the database
type User struct {
	ID       int
	Name     string
	Email    string
}
