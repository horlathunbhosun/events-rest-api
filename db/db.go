package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/horlathunbhosun/events-rest-api/internal/utility"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
	godotenv.Load()
	connStr, dbServer := utility.ConnectionStringAndDriver()
	var err error

	DB, err = sql.Open(dbServer, connStr)

	if err != nil {
		panic("Could not connect to the database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		email TEXT NOT NULL,
		password TEXT NOT NULL,
		datecreated DATETIME
	)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		log.Fatalln(err)
		panic("Can not users  table")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events(
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		log.Fatalln(err)
		panic("Can not event table")
	}
}
