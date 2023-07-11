package main

import (
	"github.com/fschmtt/thw-leaderboard/api"
	"github.com/fschmtt/thw-leaderboard/db"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	d, err := db.Connect()
	if err != nil {
		panic(err.Error())
	}

	driver, err := mysql.WithInstance(d, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)
	if err != nil {
		panic(err.Error())
	}

	err = m.Up()
	if err != nil {
		panic(err.Error())
	}

	err = api.Start(d)
	if err != nil {
		panic(err.Error())
	}
}
