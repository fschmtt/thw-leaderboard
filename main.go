package main

import (
	"github.com/fschmtt/thw-leaderboard/api"
)

func main() {
	db, err := api.Connect()
	if err != nil {
		panic(err.Error())
	}

	err = api.Start(db)
	if err != nil {
		panic(err.Error())
	}
}
