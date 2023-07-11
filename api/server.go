package api

import (
	"database/sql"
	"encoding/json"
	"github.com/fschmtt/thw-leaderboard/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

type API struct {
	Db *sql.DB
}

func Start(db *sql.DB) error {
	api := &API{
		Db: db,
	}

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/api/competitor", api.getAllCompetitors)
	r.POST("/api/competitor", api.addNewCompetitor)

	err := r.Run(os.Getenv("SERVER_ADDRESS"))
	if err != nil {
		return err
	}

	return nil
}

func (api *API) getAllCompetitors(c *gin.Context) {
	competitors, err := db.GetCompetitors(api.Db)
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"competitors": competitors,
	})
}

func (api *API) addNewCompetitor(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		panic(err.Error())
	}

	var competitor db.NewCompetitor
	err = json.Unmarshal(body, &competitor)
	if err != nil {
		panic(err.Error())
	}

	err = db.AddNewCompetitor(competitor, api.Db)
	if err != nil {
		panic(err.Error())
	}

	c.Status(http.StatusCreated)
}
