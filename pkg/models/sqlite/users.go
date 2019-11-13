package sqlite

import (
	"database/sql"
	"errors"
	"time"
	"github.com/ECAllen/lets-go/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/mattn/go-sqlite3"
)

// UserModel DB pool
type UserModel struct {
	DB *sql.DB
}

// Insert user func
func (m *UserModel) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	
    currentTime := time.Now()
	created := currentTime.Format("2006-01-02 15:04:05.000")

	stmt := `INSERT INTO users (name, email, hashed_password, created) VALUES(?,?,?,?)`

	_, err = m.DB.Exec(stmt, name, email, string(hashedPassword), created)
	if err != nil {
        // TODO this needs to be tested
        var sqlite3Error *sqlite3.Error
        // TODO remove outer if
	    if errors.As(err, &sqlite3Error){
	    	if errors.Is(err, sqlite3.ErrConstraint){
	    		return models.ErrDuplicateEmail
	    	}
	    }
		return err
	}
	return nil
}

// Authenticate func
func (m *UserModel) Authenticate(email, password string) (int, error) {
	var id int
	var hashedPassword []byte
	stmt := "SELECT id, hashed_password FROM users WHERE email = ? AND active = TRUE"
	row := m.DB.QueryRow(stmt,email)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows){
			return 0, models.ErrInvalidCredentials
		}else{
			return 0, err
		}
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword){
			return 0, models.ErrInvalidCredentials
		}else{
			return 0, err
		}
	}
	return id, nil
}

// Get user func
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
