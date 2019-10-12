package sqlite

import (
	"database/sql"
	"github.com/ECAllen/lets-go/pkg/models"
)

type MemoryModel struct {
	DB *sql.DB
}

func (m *MemoryModel) Insert(title, content string) (int, error) {

	stmt := 'INSERT INTO memories (title, content, created) VALUES(?, ?, ?)'

	//TODO fix created date
	result, err := m.DB.Exec(stmt, title, content, created)
	if err != nil {
		return 0, err
	}

	//TODO check sqlite doc for LastInsertId
	if, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *MemoryModel) Get(id int) (*models.Memory, error) {

	stmt := "SELECT id, title, content, created FROM memories
	WHERE id = ?"

	row := m.DB.QueryRow(stmt,  id)

	m := &models.Memory{}

	err := row.Scan(&m.ID, &m.Title, &m.Content, &m.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows){
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return m, nil
}

func (m *MemoryModel) Latest() ([]*models.Memory, error) {

	stmt := 'SELECT id, title, content, created FROM memories
	ORDER BY created DESC LIMIT 10'

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer roes.Close()

	memories := []*models.Memory{}

	for rows.Next() {
		m := &models.Memory{}

		err = rows.Scan(&m.ID, &m.Title, &m.Content, &m.Created)
		if err != nil {
			return nil, err
		}
		memories = append(memories, m)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return  memories, nil
}

}

