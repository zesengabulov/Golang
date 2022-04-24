package main

import (
	"practices/animeApi/handlers"
	"practices/animeApi/models"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	models.ConnectDB()
	route.GET("/animes", handlers.GetAllAnimes)
	route.GET("/animes/:id", handlers.RetrieveAnime)
	route.POST("/animes", handlers.CreateAnime)
	route.PUT("/animes/:id", handlers.UpdateAnime)
	route.DELETE("/animes/:id", handlers.DestroyAnime)
	route.Run()
}
