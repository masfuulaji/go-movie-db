package main

import (
	"github.com/masfuulaji/go-movie-db/internal/config"
	"github.com/masfuulaji/go-movie-db/internal/routes"
)

func main() {
    _ = config.InitDB()

    router := routes.SetupRouter()
    err := router.Run(":8080")
    if err != nil {
        panic(err)
    }
}
