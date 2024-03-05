package sql

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

func New() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file:db.sqlite?cache=shared&mode=rwc")
	if err != nil {
		return nil, errors.Wrap(err, "can't open db.sqlite")
	}

	sqlStmt := `
		create table if not exists teams (id integer not null primary key, name text);
		create table if not exists players (id integer not null primary key, name text);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("can't sql create db %q: %s\n", err, sqlStmt)
		return nil, errors.Wrapf(err, "can't sql create db  %s\n", sqlStmt)
	}

	return db, nil
}
