package database

import (
	"datestore/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type DateDB interface {
	Open() error
	Close() error
	PostYear(p *models.Date) (int64, error)
	GetYear() ([]*models.Date, error)
}

type DB struct {
	db *sqlx.DB
}

func (d *DB) Open() error {
	pg, err := sqlx.Open("postgres", pgConnStr)
	if err != nil {
		log.Println("DB connection Error")
		return err
	}
	log.Println("DB connection successful!")
	pg.MustExec(createSchema)
	d.db = pg
	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}
