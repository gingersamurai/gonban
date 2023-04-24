package storage

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"gonban/internal/config"
	"gonban/internal/entity"
)

type PostgresTask struct {
	Id          int64
	Status      sql.NullString
	Name        sql.NullString
	Description sql.NullString
	Performer   sql.NullString
	Deadline    sql.NullTime
}

func (p PostgresTask) Cast() entity.Task {
	r := entity.Task{
		Id:          int(p.Id),
		Status:      p.Status.String,
		Name:        p.Name.String,
		Description: p.Description.String,
		Performer:   p.Performer.String,
		Deadline:    p.Deadline.Time,
	}
	return r
}

type PostgresTaskStorage struct {
	conn *sqlx.DB
}

func NewPostgresTaskStorage(cfg config.PostgresConfig) (*PostgresTaskStorage, error) {
	connectionInfo := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.SSLMode, cfg.Password,
	)
	conn, err := sqlx.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	if err = conn.Ping(); err != nil {
		return nil, err
	}
	return &PostgresTaskStorage{conn: conn}, nil
}

func (ps *PostgresTaskStorage) Add(task entity.Task) (int, error) {
	dbRequest := `INSERT INTO tasks(status, name, description, performer, deadline) VALUES (:status, :name, :description, :performer, :deadline) RETURNING id`
	row, err := ps.conn.NamedQuery(dbRequest, task)
	if err != nil {
		return 0, fmt.Errorf("MemoryStorage.Add(): %w", err)
	}
	row.Next()
	var id int64
	err = row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("MemoryStorage.Add(): %w", err)
	}
	return int(id), nil
}

func (ps *PostgresTaskStorage) GetById(id int) (entity.Task, error) {
	dbRequest := `SELECT * FROM tasks WHERE id = $1`
	postgresR := PostgresTask{}
	row := ps.conn.QueryRowx(dbRequest, id)
	if err := row.StructScan(&postgresR); err != nil {
		return entity.Task{}, err
	}
	result := postgresR.Cast()
	return result, nil
}

func (ps *PostgresTaskStorage) GetAll() []entity.Task {
	dbRequest := `SELECT * FROM tasks`
	rows, err := ps.conn.Queryx(dbRequest)
	if err != nil {
		panic(err)
	}
	var result []entity.Task
	postgresR := PostgresTask{}
	for rows.Next() {
		if err := rows.StructScan(&postgresR); err != nil {
			panic(err)
		}
		result = append(result, postgresR.Cast())
	}
	return result
}

func (ps *PostgresTaskStorage) DeleteById(id int) error {
	dbRequest := `DELETE FROM tasks WHERE id = $1`
	res, err := ps.conn.Exec(dbRequest, id)
	if err != nil {
		return err
	}
	deleted, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if deleted == 0 {
		return errors.New(fmt.Sprintf("task with id %v not found", id))
	}
	return err
}
