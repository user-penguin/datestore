package database

const createSchema = `
CREATE TABLE IF NOT EXISTS date
(
	id SERIAL PRIMARY KEY,
	year TEXT,
	curFrom TEXT,
	curTo TEXT,
	ratio TEXT
)
`

var insertYearSchema = `
INSERT INTO year(year, curFrom, curTo, ratio) VALUES ($1, $2, $3, $4) RETURNING id
`
