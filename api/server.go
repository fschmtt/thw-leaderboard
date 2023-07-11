package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Start() {
	var port, present = os.LookupEnv("SERVER_ADDRESS")
	if present != true {
		panic("SERVER_ADDRESS is not set")
	}

	r := gin.Default()
	r.GET("/api/all", getAllCompetitors)
	r.GET("/api/top-3", getTop3Competitors)
	r.GET("/api/last-5", getLast5Competitors)

	err := r.Run(port)
	if err != nil {
		panic(err.Error())
	}
}

func getAllCompetitors(c *gin.Context) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	competitors, err := GetAllCompetitors(db)
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"competitors": competitors,
	})
}

func getTop3Competitors(c *gin.Context) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	competitors, err := GetTop3Competitors(db)
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"competitors": competitors,
	})
}

func getLast5Competitors(c *gin.Context) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	competitors, err := GetLast5Competitors(db)
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"competitors": competitors,
	})
}
