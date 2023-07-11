package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	"math"
	"math/rand"
	"os"
)

type NewCompetitor struct {
	Name    string `json:"name" validate:"required"`
	OffsetX int    `json:"offsetX" validate:"required"`
	OffsetY int    `json:"offsetY" validate:"required"`
}

type Competitor struct {
	Id          int    `json:"id"`
	Identifier  string `json:"identifier"`
	Name        string `json:"name"`
	OffsetX     int    `json:"offsetX"`
	OffsetY     int    `json:"offsetY"`
	FaultPoints int    `json:"faultPoints"`
	CreatedAt   string `json:"createdAt"`
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

func GetAllCompetitors(db *sql.DB) ([]Competitor, error) {
	rows, err := db.Query("SELECT id, identifier, name, offset_x, offset_y, fault_points, created_at FROM competitor ORDER BY fault_points ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var competitors []Competitor
	for rows.Next() {
		var competitor Competitor
		err := rows.Scan(&competitor.Id, &competitor.Identifier, &competitor.Name, &competitor.OffsetX, &competitor.OffsetY, &competitor.FaultPoints, &competitor.CreatedAt)
		if err != nil {
			return nil, err
		}

		competitors = append(competitors, competitor)
	}

	return competitors, nil
}

func GetTop3Competitors(db *sql.DB) ([]Competitor, error) {
	rows, err := db.Query("SELECT id, identifier, name, offset_x, offset_y, fault_points, created_at FROM competitor ORDER BY fault_points ASC LIMIT 3")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var competitors []Competitor
	for rows.Next() {
		var competitor Competitor
		err := rows.Scan(&competitor.Id, &competitor.Identifier, &competitor.Name, &competitor.OffsetX, &competitor.OffsetY, &competitor.FaultPoints, &competitor.CreatedAt)
		if err != nil {
			return nil, err
		}

		competitors = append(competitors, competitor)
	}

	return competitors, nil
}

func GetLatest5Competitors(db *sql.DB) ([]Competitor, error) {
	rows, err := db.Query("SELECT id, identifier, name, offset_x, offset_y, fault_points, created_at FROM competitor ORDER BY id DESC LIMIT 3")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var competitors []Competitor
	for rows.Next() {
		var competitor Competitor
		err := rows.Scan(&competitor.Id, &competitor.Identifier, &competitor.Name, &competitor.OffsetX, &competitor.OffsetY, &competitor.FaultPoints, &competitor.CreatedAt)
		if err != nil {
			return nil, err
		}

		competitors = append(competitors, competitor)
	}

	return competitors, nil
}

func AddNewCompetitor(nc NewCompetitor, db *sql.DB) error {
	x := float64(nc.OffsetX)
	y := float64(nc.OffsetY)
	faultPoints := math.Sqrt(math.Pow(x, 2) + (math.Pow(y, 2)))
	log.Println(faultPoints)

	stmt, err := db.Prepare("INSERT INTO competitor (identifier, name, offset_x, offset_y, fault_points) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(rand.Int(), nc.Name, nc.OffsetX, nc.OffsetY, faultPoints)
	if err != nil {
		return err
	}

	return nil
}
