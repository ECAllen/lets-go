package models

import (
	"errors"
	"time"
)

var (
	// ErrNoRecord no matching records
	ErrNoRecord = errors.New("models: no matching record found")
	// ErrInvalidCredentials failed db checks
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	// ErrDuplicateEmail failed db contraint on email
	ErrDuplicateEmail = errors.New("models: duplicate email")
)

// Memory struct
type Memory struct {
	ID      int
	Title   string
	Content string
	Created time.Time
}

// User type
type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
	Active         bool
}
