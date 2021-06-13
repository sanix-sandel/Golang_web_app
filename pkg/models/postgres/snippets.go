package postgres

import (
	"database/sql"

	"sanix.net/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content string) (int, error) {

	//stmt := `INSERT INTO snippets (title, content) VALUES (?, ?)`
	//result, err := m.DB.Exec(stmt, title, content)
	stmt, _ := m.DB.Prepare("INSERT INTO snippets (title, content) VALUES (?, ?)")
	result, err := stmt.Exec(title, content)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId() //RowAffected returns the number of rows
	//affected by statement
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
