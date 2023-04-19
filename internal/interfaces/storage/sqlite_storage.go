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
	return entity.Task{}, nil
}

func (ss *SqliteTaskStorage) GetAll() []entity.Task {
	dbRequest := `SELECT id, status, name, description, performer, deadline from tasks`
	deadlineLayout := time.RFC3339

	rows, _ := ss.conn.Query(dbRequest)
	result := []entity.Task{}
	for rows.Next() {
		sqliteR := SqliteTask{}
		err := rows.Scan(&sqliteR.Id, &sqliteR.Status, &sqliteR.Name, &sqliteR.Description, &sqliteR.Performer, &sqliteR.Deadline)
		if err != nil {
			panic(err)
		}
		r := entity.Task{}
		r.Id = int(sqliteR.Id)
		if sqliteR.Status.Valid {
			r.Status = sqliteR.Status.String
		}
		if sqliteR.Name.Valid {
			r.Name = sqliteR.Name.String
		}
		if sqliteR.Description.Valid {
			r.Description = sqliteR.Description.String
		}
		if sqliteR.Performer.Valid {
			r.Performer = sqliteR.Performer.String
		}
		if sqliteR.Deadline.Valid {
			r.Deadline, err = time.Parse(deadlineLayout, sqliteR.Deadline.String)
			if err != nil {
				panic(err)
			}
		}

		result = append(result, r)
	}
	return result
}

func (ss *SqliteTaskStorage) DeleteById(id int) error {
	return nil
}

func (ss *SqliteTaskStorage) Close() error {
	return nil
}
