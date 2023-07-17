package api

import (
	"database/sql"
	"encoding/json"
	"github.com/fschmtt/thw-leaderboard/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

	r.Use(static.Serve("/", static.LocalFile("./frontend", false)))
	r.Use(static.Serve("/assets", static.LocalFile("./frontend/assets", false)))

	r.GET("/api/competitor", api.getAllCompetitors)
	r.POST("/api/competitor", api.addNewCompetitor)
	r.GET("/api/competitor/count", api.countAllCompetitors)

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

	if competitors == nil {
		competitors = []db.Competitor{}
	}

	c.JSON(http.StatusOK, gin.H{
		"competitors": competitors,
	})
}

func (api *API) countAllCompetitors(c *gin.Context) {
	competitors, err := db.CountCompetitors(api.Db)
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

	var nc db.NewCompetitor
	err = json.Unmarshal(body, &nc)
	if err != nil {
		panic(err.Error())
	}

	validate := validator.New()
	err = validate.Struct(nc)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	err = db.AddNewCompetitor(nc, api.Db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.Status(http.StatusCreated)
}
