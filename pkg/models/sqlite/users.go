package sqlite

import (
	"database/sql"

	"github.com/ECAllen/lets-go/pkg/models"
	"golang.org/x/crypto/bcrypt"
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

	stmt := `INSERT INTO users (name, email, hashed_password, created) VALUES(?,?,?, UTC_TIMEPSTAMP())`

	_, err = m.DB.Exec(stmt, name, email, string(hashedPassword))
	if err != nil {
		/* TODO this needs to be changed
		var sqliteError *sqlite.Error
		if errors.As(err, &sqliteError) {
			if sqliteError.Code == 1062 && strings.Contains(mySQLiteError.Message, "users_uc_email") {
				return models.ErrDuplicateEmail
			}
		}
		*/
		return err
	}
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
