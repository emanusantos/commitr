package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func Init() {
	dbName := os.Getenv("DATABASE_NAME")

	db, err := sql.Open("sqlite3", "./"+dbName)

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	createTables(db)
}

func createTables(db *sql.DB) {
	createTables := `
	create table if not exists user (
    id uuid primary key, 
    username text not null,
    token text not null,
    refresh_token text
  );

  create table if not exists repository (
    id uuid not null primary key,
    name text not null,
    organization text not null,
    user_id uuid not null,
    foreign key (user_id) references user (id)
  );
	`

	_, err := db.Exec(createTables)

	if err != nil {
		log.Fatal(err)
	}

	db.Close()
}
