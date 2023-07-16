package main

import (
	"github.com/fschmtt/thw-leaderboard/db"
	"github.com/fschmtt/thw-leaderboard/server"
)

func main() {
	d, err := db.Connect()
	if err != nil {
		panic(err.Error())
	}

	err = server.Start(d)
	if err != nil {
		panic(err.Error())
	}
}
