package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port           string
	GithubClientId string
	GithubSecret   string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}
	return &Config{
		Port:           getEnv("PORT", "3000"),
		GithubClientId: getEnv("GITHUB_CLIENT_ID", ""),
		GithubSecret:   getEnv("GITHUB_SECRET", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
