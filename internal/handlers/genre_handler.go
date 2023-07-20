package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/masfuulaji/go-movie-db/internal/models"
	"github.com/masfuulaji/go-movie-db/internal/repositories"
)

func CreateGenreHandler(c *gin.Context) {
    var genre models.Genre
    err := c.BindJSON(&genre)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    err = repositories.CreateGenre(genre)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{"message": "Genre created"})
}

func GetGenreHandler(c *gin.Context) {
    genreID := c.Param("genre_id")

    genre, err := repositories.GetGenreById(genreID)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, genre)
}

func GetAllGenresHandler(c *gin.Context) {
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

    genres, err := repositories.GetAllGenres(page, limit)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, genres)
}

func UpdateGenreHandler(c *gin.Context) {
    var genre models.Genre
    genreID := c.Param("genre_id")

    err := c.BindJSON(&genre)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    result, err := repositories.UpdateGenre(genreID, genre)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, result)
}

func DeleteGenreHandler(c *gin.Context) {
    genreID := c.Param("genre_id")

    err := repositories.DeleteGenre(genreID)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{"message": "Genre deleted"})
}
