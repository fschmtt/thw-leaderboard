package main

import (
	"github.com/fschmtt/thw-leaderboard/api"
	"github.com/fschmtt/thw-leaderboard/db"
)

func main() {
	d, err := db.Connect()
	if err != nil {
		panic(err.Error())
	}

	err = api.Start(d)
	if err != nil {
		panic(err.Error())
	}
}
