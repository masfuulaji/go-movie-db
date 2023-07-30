package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/masfuulaji/go-movie-db/internal/models"
	"github.com/masfuulaji/go-movie-db/internal/repositories"
)

type MovieHandler struct {
    repo repositories.MovieRepository
    Validate *validator.Validate
}

func NewMovieHandler(repo repositories.MovieRepository, validate *validator.Validate) *MovieHandler {
    return &MovieHandler{repo: repo, Validate: validate}
}

func (h *MovieHandler) CreateMovieHandler(c *gin.Context) {
    var movie models.Movie
    err := c.BindJSON(&movie)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err = h.repo.CreateMovie(movie)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Movie created"})
}

func (h *MovieHandler) GetMovieHandler(c *gin.Context) {
    movieID := c.Param("movie_id")
    movie, err := h.repo.GetMovieById(movieID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, movie)
}

func (h *MovieHandler) GetAllMoviesHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    movies, err := h.repo.GetAllMovies(page, limit)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, movies)
}

func (h *MovieHandler) UpdateMovieHandler(c *gin.Context) {
    var movie models.Movie
    movieID := c.Param("movie_id")

    err := c.BindJSON(&movie)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result, err := h.repo.UpdateMovie(movieID, movie)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, result)
}

func (h *MovieHandler) DeleteMovieHandler(c *gin.Context) {
    movieID := c.Param("movie_id")
    err := h.repo.DeleteMovie(movieID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Movie deleted"})
}
