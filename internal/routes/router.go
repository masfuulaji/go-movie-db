package routes

import (
	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/masfuulaji/go-movie-db/internal/config"
	"github.com/masfuulaji/go-movie-db/internal/handlers"
	"github.com/masfuulaji/go-movie-db/internal/helper"
	"github.com/masfuulaji/go-movie-db/internal/middleware"
	"github.com/masfuulaji/go-movie-db/internal/repositories"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(CORS())
	db := config.InitDB()

	validator := validator.New()

	router.Use(helper.ErrorHandler())

	genreHandler := handlers.NewGenreHandler(repositories.NewGenreRepository(db), validator)
	genre := router.Group("/genre")
	{
		genre.GET("", genreHandler.GetAllGenresHandler)
		genre.POST("", genreHandler.CreateGenreHandler)
		genre.GET("/:genreID", genreHandler.GetGenreHandler)
		genre.PUT("/:genreID", genreHandler.UpdateGenreHandler)
		genre.DELETE("/:genreID", genreHandler.DeleteGenreHandler)
	}

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

	userHandler := handlers.NewUserHandler(repositories.NewUserRepository(db), validator)
	user := router.Group("/user")
	{
		user.GET("/", userHandler.GetAllUserHandler)
		user.POST("/", userHandler.CreateUserHandler)
		user.GET("/:userID", userHandler.GetUserHandler)
		user.PUT("/:userID", userHandler.UpdateUserHandler)
		user.DELETE("/:userID", userHandler.DeleteUserHandler)
	}

	collectionHandler := handlers.NewCollectionHandler(repositories.NewCollectionRepository(db), validator)
	collection := router.Group("/collection")
	{
		collection.GET("/", collectionHandler.GetAllCollectionsHandler)
		collection.POST("/", collectionHandler.CreateCollectionHandler)
		collection.GET("/:collectionID", collectionHandler.GetCollectionHandler)
		collection.PUT("/:collectionID", collectionHandler.UpdateCollectionHandler)
		collection.DELETE("/:collectionID", collectionHandler.DeleteCollectionHandler)
	}

	movieHandler := handlers.NewMovieHandler(repositories.NewMovieRepository(db), validator)
	movie := router.Group("/movie")
	{
		movie.GET("/", movieHandler.GetAllMoviesHandler)
		movie.POST("/", movieHandler.CreateMovieHandler)
		movie.GET("/:movieID", movieHandler.GetMovieHandler)
		movie.PUT("/:movieID", movieHandler.UpdateMovieHandler)
		movie.DELETE("/:movieID", movieHandler.DeleteMovieHandler)
	}

	return router
}
