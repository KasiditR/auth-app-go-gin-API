package main

import (
	"github.com/KasiditR/auth-app-go-gin-API/config"
	"github.com/KasiditR/auth-app-go-gin-API/internal/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg := config.LoadConfig()
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization"},
	}))

	router.Use((gin.Logger()))

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})

	routes.AuthRoutes(router)

	log.Fatal(router.Run(":" + cfg.Port))
}
