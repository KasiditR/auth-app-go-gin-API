package handlers

import (
	"github.com/KasiditR/auth-app-go-gin-API/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetGoogleUserData() gin.HandlerFunc {
	return func(c *gin.Context) {
		idToken := c.Param("id_token")
		userData, err := services.GetGoogleUserData(idToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, userData)
	}
}

func GetGithubUserData() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Param("code")
		accessToken, err := services.GetGithubAccessToken(code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		userData, err := services.GetGithubUserData(accessToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, userData)
	}
}
