package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/masfuulaji/go-movie-db/internal/config"
	"github.com/masfuulaji/go-movie-db/internal/routes"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        panic(err)
    }
    _ = config.InitDB()

    router := routes.SetupRouter()

    port := os.Getenv("PORT")
    err = router.Run(":" + port)
    if err != nil {
        panic(err)
    }
}
