package models

type Date struct {
	ID      int64  `db:"id"`
	Year    string `db:"year"`
	CurFrom string `db:"curFrom"`
	CurTo   string `db:"curTo"`
	Ratio   string `db:"ratio"`
}

type DateRequest struct {
	Year    string `db:"year"`
	CurFrom string `db:"curFrom"`
	CurTo   string `db:"curTo"`
	Ratio   string `db:"ratio"`
}
