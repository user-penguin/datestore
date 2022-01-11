package database

const createSchema = `
CREATE TABLE IF NOT EXISTS dates
(
	id SERIAL PRIMARY KEY,
	year TEXT,
	curFrom TEXT,
	curTo TEXT,
	ratio TEXT
)
`
