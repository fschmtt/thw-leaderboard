package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"math"
	"os"
)

type NewCompetitor struct {
	Name        *string `json:"name"`
	Identifier  int     `json:"identifier" validate:"required"`
	Target      float64 `json:"target" validate:"required"`
	Measurement float64 `json:"measurement" validate:"required"`
}

type Competitor struct {
	Id          int     `json:"id"`
	Name        *string `json:"name"`
	Identifier  int     `json:"identifier"`
	Target      float64 `json:"target" validate:"required"`
	Measurement float64 `json:"measurement"`
	Offset      float64 `json:"offset"`
	CreatedAt   string  `json:"createdAt"`
	Rank        int     `json:"rank"`
}

func Connect() (*sql.DB, error) {
	cfg := mysql.Config{
		User:   os.Getenv("DATABASE_USER"),
		Passwd: os.Getenv("DATABASE_PASSWORD"),
		Net:    "tcp",
		Addr:   os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT"),
		DBName: os.Getenv("DATABASE_NAME"),
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetCompetitors(db *sql.DB) ([]Competitor, error) {
	rows, err := db.Query("SELECT id, name, identifier, target, measurement, offset, created_at FROM competitor ORDER BY offset ASC, id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var competitors []Competitor
	var rank = 1
	for rows.Next() {
		var competitor Competitor
		err := rows.Scan(&competitor.Id, &competitor.Name, &competitor.Identifier, &competitor.Target, &competitor.Measurement, &competitor.Offset, &competitor.CreatedAt)
		if err != nil {
			return nil, err
		}

		competitor.Rank = rank
		competitor.Offset = metersToCentimeters(competitor.Offset)
		competitors = append(competitors, competitor)

		rank++
	}

	return competitors, nil
}

func CountCompetitors(db *sql.DB) (int, error) {
	rows, err := db.Query("SELECT COUNT(id) FROM competitor")
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, err
		}
	}

	return count, nil
}

func AddNewCompetitor(nc NewCompetitor, db *sql.DB) error {
	target := float64(nc.Target)
	measurement := float64(nc.Measurement)
	offset := math.Abs(target - measurement)

	stmt, err := db.Prepare("INSERT INTO competitor (name, identifier, target, measurement, offset) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(nc.Name, nc.Identifier, nc.Target, nc.Measurement, offset)
	if err != nil {
		return err
	}

	return nil
}

func metersToCentimeters(val float64) float64 {
	return val * 100
}
