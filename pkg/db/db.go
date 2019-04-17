package db

import (
	"database/sql"
	"log"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) Init(dbType string, connectionCof string) error {
	db, err := sql.Open(dbType, connectionCof)
	if err != nil {
		return err
	}
	r.DB = db
	log.Println("DB in init: ", r.DB)
	return nil
}
