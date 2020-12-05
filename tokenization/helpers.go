package tokenization

import "github.com/jmoiron/sqlx"

func opendDb(filename string) (*sqlx.DB, error) {
	if db, err := sqlx.Connect("sqlite3", filename); err != nil {
		return nil, err
	} else {
		return db, nil
	}
}
