package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/masfuulaji/go-movie-db/internal/models"
	"github.com/masfuulaji/go-movie-db/internal/repositories"
)

type GenreHandler struct {
    repo repositories.GenreRepository
    Validate *validator.Validate
}

func NewGenreHandler(repo repositories.GenreRepository, validate *validator.Validate) *GenreHandler {
    return &GenreHandler{repo: repo, Validate: validate}
}

func (h *GenreHandler) CreateGenreHandler(c *gin.Context) {
    var genre models.Genre
    err := c.BindJSON(&genre)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    err = h.Validate.Struct(genre)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    err = h.repo.CreateGenre(genre)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{"message": "Genre created"})
}

func (h *GenreHandler) GetGenreHandler(c *gin.Context) {
    genreID := c.Param("genre_id")

    genre, err := h.repo.GetGenreById(genreID)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, genre)
}

func (h *GenreHandler) GetAllGenresHandler(c *gin.Context) {
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

    genres, err := h.repo.GetAllGenres(page, limit)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, genres)
}

func (h *GenreHandler) UpdateGenreHandler(c *gin.Context) {
    var genre models.Genre
    genreID := c.Param("genre_id")

    err := c.BindJSON(&genre)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    err = h.Validate.Struct(genre)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    result, err := h.repo.UpdateGenre(genreID, genre)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, result)
}

func (h *GenreHandler) DeleteGenreHandler(c *gin.Context) {
    genreID := c.Param("genre_id")

    err := h.repo.DeleteGenre(genreID)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{"message": "Genre deleted"})
}
