package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masfuulaji/go-movie-db/internal/models"
	"github.com/masfuulaji/go-movie-db/internal/repositories"
)

func CreateMovieHandler(c *gin.Context) {
    var movie models.Movie
    err := c.BindJSON(&movie)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err = repositories.CreateMovie(movie)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Movie created"})
}

func GetMovieHandler(c *gin.Context) {
    movieID := c.Param("movie_id")
    movie, err := repositories.GetMovieById(movieID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, movie)
}

func GetAllMoviesHandler(c *gin.Context) {
    movies, err := repositories.GetAllMovies()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, movies)
}

func UpdateMovieHandler(c *gin.Context) {
    var movie models.Movie
    movieID := c.Param("movie_id")

    err := c.BindJSON(&movie)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result, err := repositories.UpdateMovie(movieID, movie)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, result)
}

func DeleteMovieHandler(c *gin.Context) {
    movieID := c.Param("movie_id")
    err := repositories.DeleteMovie(movieID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Movie deleted"})
}
