package repositories

import (
	"crudPackages/models"
	"database/sql"
	"errors"

	// driver implementation
	_ "github.com/go-sql-driver/mysql"
)

//DB : data base implementation
type DB struct {
	*sql.DB
}

// NewDB : Init data access
func NewDB(dataSourceName string) (*DB, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

// GetSubject : Get a subject by id
func (db DB) GetSubject(id int) (models.Subject, error) {
	rows, err := db.Query("SELECT id, name, professor FROM subjects WHERE id = ?", id)
	var s models.Subject
	if err != nil {
		return s, err
	}

	if rows.Next() {
		err = rows.Scan(&s.ID, &s.Name, &s.Professor)
		if err != nil {
			return s, err
		}
	}
	return s, nil
}

//AddSubject : Add a new subject
func (db DB) AddSubject(s models.Subject) (models.Subject, error) {
	stm, err := db.Prepare("INSERT INTO subjects (name, professor) VALUES (?, ?)")
	if err != nil {
		return s, err
	}
	result, err := stm.Exec(s.Name, s.Professor)
	if err != nil {
		return s, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return s, err
	}
	if affected != 1 {
		return s, errors.New("More than one record affected at AddSubject()")
	}
	last, err := result.LastInsertId()
	if err != nil {
		return s, err
	}
	return db.GetSubject(int(last))
}
