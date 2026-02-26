package storage

import (
	"context"
	"database/sql"

	"github.com/northmaxota/task-api/internal/task"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{
		db: db,
	}
}

func (s *PostgresStorage) Create(ctx context.Context, t task.Task) (task.Task, error) {
	query := `INSERT INTO tasks (title, done) VALUES ($1, $2) RETURNING id, done`
	row := s.db.QueryRowContext(ctx, query, t.Title, t.Done)
	var id int
	var done bool
	if err := row.Scan(&id, &done); err != nil {
		return task.Task{}, err
	}
	t.ID = id
	t.Done = done
	return t, nil
}

func (s *PostgresStorage) GetAll(ctx context.Context) ([]task.Task, error) {
	query := `SELECT id, title, done FROM tasks`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []task.Task
	for rows.Next() {
		var t task.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Done); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}
