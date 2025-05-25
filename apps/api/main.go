package main

import (
	"log"
	"os"
	"yoga/api/db"
	"yoga/api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    godotenv.Load()

    _, err := db.ConnectDB()
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    router := gin.Default()
    routes.RegisterRoutes(router)

    port := os.Getenv("PORT")
    if port == "" {
        port = "3001"
    }

    log.Printf("API running on port %s", port)
    router.Run(":" + port)
}
