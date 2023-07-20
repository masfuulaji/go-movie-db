package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/masfuulaji/go-movie-db/internal/handlers"
	"github.com/masfuulaji/go-movie-db/internal/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	home := router.Group("/")
	{
		home.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello World",
			})
		})
	}

	login := router.Group("/auth")
	{
		login.POST("/login", handlers.LoginHandler)
	}

	logout := router.Group("/auth")
	logout.Use(middleware.AuthMiddleware())
	{
		logout.POST("/logout", handlers.LogoutHandler)
	}

	user := router.Group("/user")
	// user.Use(middleware.AuthMiddleware())
	{
		user.GET("/", handlers.GetAllUserHandler)
		user.POST("/", handlers.CreateUserHandler)
		user.GET("/:userID", handlers.GetUserHandler)
		user.PUT("/:userID", handlers.UpdateUserHandler)
		user.DELETE("/:userID", handlers.DeleteUserHandler)
	}

    collection := router.Group("/collection")
    {
        collection.GET("/", handlers.GetAllCollectionsHandler)
        collection.POST("/", handlers.CreateCollectionHandler)
        collection.GET("/:collectionID", handlers.GetCollectionHandler)
        collection.PUT("/:collectionID", handlers.UpdateCollectionHandler)
        collection.DELETE("/:collectionID", handlers.DeleteCollectionHandler)
    }

    movie := router.Group("/movie")
    {
        movie.GET("/", handlers.GetAllMoviesHandler)
        movie.POST("/", handlers.CreateMovieHandler)
        movie.GET("/:movieID", handlers.GetMovieHandler)
        movie.PUT("/:movieID", handlers.UpdateMovieHandler)
        movie.DELETE("/:movieID", handlers.DeleteMovieHandler)
    }

    genre := router.Group("/genre")
    {
        genre.GET("/", handlers.GetAllGenresHandler)
        genre.POST("/", handlers.CreateGenreHandler)
        genre.GET("/:genreID", handlers.GetGenreHandler)
        genre.PUT("/:genreID", handlers.UpdateGenreHandler)
        genre.DELETE("/:genreID", handlers.DeleteGenreHandler)
    }

	return router
}
