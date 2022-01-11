package database

import "datestore/models"

func (d *DB) PostYear(p *models.Date) error {
	res, err := d.db.Exec(insertYearSchema, p.Year, p.CurFrom, p.CurTo, p.Ratio)
	if err != nil {
		return err
	}
	// TODO по посту возвращать id новой записи
	_, err = res.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) GetYear() ([]*models.Date, error) {
	var years []*models.Date
	err := d.db.Select(&years, "SELECT * FROM year")
	if err != nil {
		return nil, err
	}
	return years, nil
}
