package database

import "database/sql"

const (
	users = `
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  age INT,
  first_name TEXT,
  last_name TEXT,
  email TEXT UNIQUE NOT NULL
);
`
)

func CreateUserTable(db *sql.DB) (sql.Result, error) {
	return db.Exec(users)
}
