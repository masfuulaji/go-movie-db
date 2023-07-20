package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/masfuulaji/go-movie-db/internal/models"
	"github.com/masfuulaji/go-movie-db/internal/repositories"
)

func CreateCollectionHandler(c *gin.Context) {
	var collection models.Collection
	err := c.BindJSON(&collection)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = repositories.CreateCollection(collection)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Collection created"})
}

func GetCollectionHandler(c *gin.Context) {
	collectionID := c.Param("collection_id")

	collection, err := repositories.GetCollectionById(collectionID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, collection)
}

func GetAllCollectionsHandler(c *gin.Context) {
    collections, err := repositories.GetAllCollections()
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, collections)
}

func UpdateCollectionHandler(c *gin.Context) {
	var collection models.Collection
	collectionID := c.Param("collection_id")

	err := c.BindJSON(&collection)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

    result, err := repositories.UpdateCollection(collectionID, collection)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

    c.JSON(200, result)
}

func DeleteCollectionHandler(c *gin.Context) {
    collectionID := c.Param("collection_id")

    err := repositories.DeleteCollection(collectionID)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{"message": "Collection deleted"})
}
