package models

import (
	"time"
)

// User is ..
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Users is ..
type Users []User
