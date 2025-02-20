package routes

import (
	"github.com/KasiditR/auth-app-go-gin-API/internal/handlers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(routes *gin.Engine) {
	authRoutes := routes.Group("/auth")
	authRoutes.GET("/github/get-user-data/:code", handlers.GetGithubUserData())
	authRoutes.GET("/google/get-user-data/:id_token", handlers.GetGoogleUserData())
}
