package main

import (
	"os"

	"github.com/joho/godotenv"
)

type ApiConfig struct {
	Version     string
	Env         string
	Port        string
	CorsMethods string
	CorsOrigins string
}

var Config = ApiConfig{
	Version:     "1.0.0",
	Env:         "default",
	Port:        "8080",
	CorsMethods: "OPTIONS,GET,POST,PUT,PATCH,DELETE",
	CorsOrigins: "*",
}

func init() {
	env := getEnv("ENV", "default")
	var envFilePath string
	if env == "default" {
		return
	} else {
		envFilePath = ".env." + env
	}

	err := godotenv.Load(envFilePath)
	if err != nil {
		return
	}

	Config.Env = env
	Config.Port = getEnv("PORT", Config.Port)
	Config.CorsMethods = getEnv("ALLOW_CORS_METHODS", Config.CorsMethods)
	Config.CorsOrigins = getEnv("ALLOW_CORS_ORIGINS", Config.CorsOrigins)
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
