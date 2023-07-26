package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/masfuulaji/go-movie-db/internal/config"
	"github.com/masfuulaji/go-movie-db/internal/handlers"
	"github.com/masfuulaji/go-movie-db/internal/helper"
	"github.com/masfuulaji/go-movie-db/internal/middleware"
	"github.com/masfuulaji/go-movie-db/internal/repositories"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
    db := config.InitDB()

    router.Use(helper.ErrorHandler())

	home := router.Group("/")
	{
		home.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello World",
			})
		})
	}

    authHandler := handlers.NewAuthHandler(repositories.NewUserRepository(db))
	login := router.Group("/auth")
	{
		login.POST("/login", authHandler.LoginHandler)
	}

	logout := router.Group("/auth")
	logout.Use(middleware.AuthMiddleware())
	{
		logout.POST("/logout", authHandler.LogoutHandler)
	}


    userHandler := handlers.NewUserHandler(repositories.NewUserRepository(db))
	user := router.Group("/user")
	// user.Use(middleware.AuthMiddleware())
	{
		user.GET("/", userHandler.GetAllUserHandler)
		user.POST("/", userHandler.CreateUserHandler)
		user.GET("/:userID", userHandler.GetUserHandler)
		user.PUT("/:userID", userHandler.UpdateUserHandler)
		user.DELETE("/:userID", userHandler.DeleteUserHandler)
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
