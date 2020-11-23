package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/utkarshmani1997/talks/12Dec2020/techniques/mocking/bootstrap"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "test"
	password = "test"
	dbname   = "test"
)

var db *sql.DB

func Setup() {
	var err error
	bootstrap.Start("test", "5432")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err != nil {
			db.Close()
		}
	}()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	_, err = CreateUserTable(db)
	if err != nil {
		panic(err)
	}
}

func TearDown() {
	db.Close()
	bootstrap.Stop("test")
}
