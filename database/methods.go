package database

import (
	"datestore/models"
	"log"
)

func (d *DB) PostYear(p *models.Date) (int64, error) {
	var id int64
	if err := d.db.QueryRow(insertYearSchema, p.Year, p.CurFrom, p.CurTo, p.Ratio).Scan(&id); err != nil {
		log.Println(err)
		return 0, err
	}
	return id, nil
}

func (d *DB) GetYear() ([]*models.Date, error) {
	var years []*models.Date
	err := d.db.Select(&years, "SELECT * FROM year")
	if err != nil {
		return nil, err
	}
	return years, nil
}
