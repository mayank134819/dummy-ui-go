package model

type User struct {
    Name     string
    Username string
    Password string
}

var Users = make(map[string]User) // In-memory DB
