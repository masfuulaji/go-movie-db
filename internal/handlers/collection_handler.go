package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/masfuulaji/go-movie-db/internal/models"
	"github.com/masfuulaji/go-movie-db/internal/repositories"
)

type CollectionHandler struct {
    repo repositories.CollectionRepository
    Validate *validator.Validate
}

func NewCollectionHandler(repo repositories.CollectionRepository, validate *validator.Validate) *CollectionHandler {
    return &CollectionHandler{repo: repo, Validate: validate}
}

func (h *CollectionHandler) CreateCollectionHandler(c *gin.Context) {
	var collection models.Collection
	err := c.BindJSON(&collection)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

    err = h.Validate.Struct(collection)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

	err = h.repo.CreateCollection(collection)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Collection created"})
}

func (h *CollectionHandler) GetCollectionHandler(c *gin.Context) {
	collectionID := c.Param("collection_id")

	collection, err := h.repo.GetCollectionById(collectionID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, collection)
}

func (h *CollectionHandler) GetAllCollectionsHandler(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    collections, err := h.repo.GetAllCollections(page, limit)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, collections)
}

func (h *CollectionHandler) UpdateCollectionHandler(c *gin.Context) {
	var collection models.Collection
	collectionID := c.Param("collection_id")

	err := c.BindJSON(&collection)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

    err = h.Validate.Struct(collection)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    result, err := h.repo.UpdateCollection(collectionID, collection)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

    c.JSON(200, result)
}

func (h *CollectionHandler) DeleteCollectionHandler(c *gin.Context) {
    collectionID := c.Param("collection_id")

    err := h.repo.DeleteCollection(collectionID)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{"message": "Collection deleted"})
}
