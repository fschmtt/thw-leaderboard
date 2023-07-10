package api

import (
	"encoding/json"
	"net/http"
	"os"
)

func Start() {
	var port, present = os.LookupEnv("SERVER_ADDRESS")
	if present != true {
		panic("SERVER_ADDRESS is not set")
	}

	http.HandleFunc("/top-3", getTop3Competitors)
	http.HandleFunc("/last-5", getLast5Competitors)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err.Error())
	}
}

func getTop3Competitors(w http.ResponseWriter, r *http.Request) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	competitors, err := GetTop3Competitors(db)
	if err != nil {
		panic(err.Error())
	}

	top3, _ := json.Marshal(competitors)

	w.WriteHeader(http.StatusOK)
	w.Write(top3)
}

func getLast5Competitors(w http.ResponseWriter, r *http.Request) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	competitors, err := GetLast5Competitors(db)
	if err != nil {
		panic(err.Error())
	}

	last5, _ := json.Marshal(competitors)

	w.WriteHeader(http.StatusOK)
	w.Write(last5)
}
