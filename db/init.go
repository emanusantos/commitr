package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func Init() {
	dbName := os.Getenv("DATABASE_NAME")

	fmt.Println(dbName)

	db, err := sql.Open("sqlite3", "./"+dbName)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	statement := `
	create table if not exists foo (id integer not null primary key, name text);
	`
	_, err = db.Exec(statement)

	if err != nil {
		log.Fatal(err)
	}
}
