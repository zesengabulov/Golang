package handlers

import (
	"fmt"
	"net/http"
	"practices/animeApi/models"

	"github.com/gin-gonic/gin"
)

type CreateAnimeInput struct {
	Name           string  `json:"name" binding:"required"`
	YearRelease    uint16  `json:"yearRelease" binding:"required"`
	Rating         float32 `json:"ratingRelease"`
	Director       string  `json:"director" binding:"required"`
	ReleaseCountry string  `json:"releaseCountry" binding:"required"`
	Genre          string  `json:"genre" binding:"required"`
}

type UpdateAnimeInput struct {
	Name           string  `json:"name"`
	YearRelease    uint16  `json:"yearRelease"`
	Rating         float32 `json:"ratingRelease"`
	Director       string  `json:"director"`
	ReleaseCountry string  `json:"releaseCountry"`
	Genre          string  `json:"genre"`
}

func GetAllAnimes(context *gin.Context) {
	var animes []models.Anime
	// models.DB.Select([]string{"id", "name", ""}).Find(&animes)
	models.DB.Find(&animes)
	context.JSON(http.StatusOK, gin.H{"allAnimes": animes})
}

func CreateAnime(context *gin.Context) {
	var input CreateAnimeInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	anime := models.Anime{
		Name:           input.Name,
		YearRelease:    input.YearRelease,
		Rating:         input.Rating,
		Director:       input.Director,
		ReleaseCountry: input.ReleaseCountry,
		Genre:          input.Genre,
	}

	models.DB.Create(&anime)
	context.JSON(http.StatusOK, gin.H{"anime": anime})
}

func UpdateAnime(context *gin.Context) {
	var anime models.Anime
	if err := models.DB.Where("id=?", context.Param("id")).First(&anime).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Записи не существует"})
		return
	}
	var input UpdateAnimeInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&anime).Update(input)
	context.JSON(http.StatusOK, gin.H{"anime": anime})
}

func RetrieveAnime(context *gin.Context) {
	var anime models.Anime
	fmt.Println(context.Param("id"))
	if err := models.DB.Where("id=?", context.Param("id")).First(&anime).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Записи не существует"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"anime": anime})
}

func DestroyAnime(context *gin.Context) {
	var anime models.Anime
	fmt.Println(context.Param("id"))
	if err := models.DB.Where("id=?", context.Param("id")).First(&anime).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Записи не существует"})
		return
	}

	models.DB.Delete(&anime)
	context.JSON(http.StatusOK, gin.H{"anime": anime})
}
