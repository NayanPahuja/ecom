package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost           string
	Port                 string
	DBUser               string
	DBPassword           string
	DBAddress            string
	DBName               string
	JWTExpirationSeconds int64
	JWTSecret            string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "admin"),
		DBPassword: getEnv("DB_PASSWORD", "admin"),
		DBAddress: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv(
			"DB_PORT", "3306")),
		DBName:               getEnv("DB_NAME", "ecom"),
		JWTExpirationSeconds: getEnvAsInt("JWT_EXPIRTION_SECONDS", 3600*24*7),
		JWTSecret:            getEnv("JWT_SECRET", "whoTheHellNamedItSecret"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		intVal, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return intVal
	}
	return fallback
}
