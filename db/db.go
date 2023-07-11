package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	"math"
	"os"
)

type NewCompetitor struct {
	Name    string `json:"name" validate:"required"`
	OffsetX int    `json:"offsetX" validate:"required"`
	OffsetY int    `json:"offsetY" validate:"required"`
}

type Competitor struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	OffsetX   int     `json:"offsetX"`
	OffsetY   int     `json:"offsetY"`
	Score     float64 `json:"score"`
	CreatedAt string  `json:"createdAt"`
	Rank      int     `json:"rank"`
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
	rows, err := db.Query("SELECT id, name, offset_x, offset_y, score, created_at FROM competitor ORDER BY score ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var competitors []Competitor
	var rank = 1
	for rows.Next() {
		var competitor Competitor
		err := rows.Scan(&competitor.Id, &competitor.Name, &competitor.OffsetX, &competitor.OffsetY, &competitor.Score, &competitor.CreatedAt)
		if err != nil {
			return nil, err
		}

		competitor.Rank = rank
		competitors = append(competitors, competitor)

		rank++
	}

	return competitors, nil
}

func AddNewCompetitor(nc NewCompetitor, db *sql.DB) error {
	x := float64(nc.OffsetX)
	y := float64(nc.OffsetY)
	score := math.Sqrt(math.Pow(x, 2) + (math.Pow(y, 2)))
	log.Println(score)

	stmt, err := db.Prepare("INSERT INTO competitor (name, offset_x, offset_y, score) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(nc.Name, nc.OffsetX, nc.OffsetY, score)
	if err != nil {
		return err
	}

	return nil
}
