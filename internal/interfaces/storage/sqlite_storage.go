package storage

import (
	"database/sql"
	"github.com/gingersamurai/gonban/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type SqliteTask struct {
	Id          int64
	Status      sql.NullString
	Name        sql.NullString
	Description sql.NullString
	Performer   sql.NullString
	Deadline    sql.NullString
}

func (s SqliteTask) Parse() entity.Task {
	r := entity.Task{
		Id:          int(s.Id),
		Status:      s.Status.String,
		Name:        s.Name.String,
		Description: s.Description.String,
		Performer:   s.Performer.String,
	}
	if s.Deadline.Valid {
		var err error
		r.Deadline, err = time.Parse(time.RFC3339, s.Deadline.String)
		if err != nil {
			panic(err)
		}
	}
	return r
}

type SqliteTaskStorage struct {
	conn *sql.DB
}

func NewSqliteTaskStorage(connectionInfo string) (*SqliteTaskStorage, error) {

	conn, err := sql.Open("sqlite3", connectionInfo)
	if err != nil {
		return nil, err
	}

	return &SqliteTaskStorage{conn: conn}, nil
}

func (ss *SqliteTaskStorage) Add(task entity.Task) int {
	dbRequest := `INSERT INTO tasks(status, name, description, performer, deadline) VALUES ($1, $2, $3, $4, $5)`
	sqliteDeadline, err := task.Deadline.MarshalText()
	if err != nil {
		panic(err)
	}
	result, _ := ss.conn.Exec(dbRequest, task.Status, task.Name, task.Description, task.Performer, sqliteDeadline)
	id, _ := result.LastInsertId()
	return int(id)
}

func (ss *SqliteTaskStorage) GetById(id int) (entity.Task, error) {
	dbRequest := `SELECT id, status, name, description, performer, deadline from tasks WHERE id = $1`
	sqliteR := SqliteTask{}
	row := ss.conn.QueryRow(dbRequest, id)
	err := row.Scan(&sqliteR.Id, &sqliteR.Status, &sqliteR.Name, &sqliteR.Description, &sqliteR.Performer, &sqliteR.Deadline)
	if err != nil {
		return entity.Task{}, err
	}
	result := sqliteR.Parse()
	return result, nil
}

func (ss *SqliteTaskStorage) GetAll() []entity.Task {
	dbRequest := `SELECT id, status, name, description, performer, deadline from tasks`

	rows, _ := ss.conn.Query(dbRequest)
	result := []entity.Task{}
	for rows.Next() {
		sqliteR := SqliteTask{}
		err := rows.Scan(&sqliteR.Id, &sqliteR.Status, &sqliteR.Name, &sqliteR.Description, &sqliteR.Performer, &sqliteR.Deadline)
		if err != nil {
			panic(err)
		}
		r := sqliteR.Parse()

		result = append(result, r)
	}
	return result
}

func (ss *SqliteTaskStorage) DeleteById(id int) error {
	dbRequest := `DELETE FROM tasks WHERE id = $1`
	_, err := ss.conn.Exec(dbRequest, id)
	return err
}

func (ss *SqliteTaskStorage) Close() error {
	return nil
}
