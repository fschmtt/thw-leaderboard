package api

import (
	"database/sql"
	"github.com/fschmtt/thw-leaderboard/db"
	"github.com/gin-gonic/gin"
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
	r.GET("/api/all", api.getAllCompetitors)
	r.GET("/api/top-3", api.getTop3Competitors)
	r.GET("/api/latest-5", api.getLatest5Competitors)

	err := r.Run(os.Getenv("SERVER_ADDRESS"))
	if err != nil {
		return err
	}

	return nil
}

func (api *API) getAllCompetitors(c *gin.Context) {
	competitors, err := db.GetAllCompetitors(api.Db)
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"competitors": competitors,
	})
}

func (api *API) getTop3Competitors(c *gin.Context) {
	competitors, err := db.GetAllCompetitors(api.Db)
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"competitors": competitors,
	})
}

func (api *API) getLatest5Competitors(c *gin.Context) {
	competitors, err := db.GetLatest5Competitors(api.Db)
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"competitors": competitors,
	})
}
