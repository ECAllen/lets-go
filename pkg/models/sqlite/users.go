package sqlite

import (
	"database/sql"

	"github.com/ECAllen/lets-go/pkg/models"
)

// UserModel DB pool
type UserModel struct {
	DB *sql.DB
}

// Insert user func
func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

// Authenticate func
func (m *UserModel) Authenticate(name, email, password string) error {
	return nil
}

// Get user func
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
