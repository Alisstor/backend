package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort int
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	JWTSecret  string
	JWTTTL     int
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	serverPort, err := strconv.Atoi(getEnv("SERVER_PORT", "8080"))
	if err != nil {
		log.Fatal("SERVER_PORT is not a valid integer")
	}

	dbPort, err := strconv.Atoi(getEnv("DB_PORT", "3306"))
	if err != nil {
		log.Fatal("DB_PORT is not a valid integer")
	}

	jwtTTL, err := strconv.Atoi(getEnv("JWT_TTL_MINUTES", "60"))
	if err != nil {
		log.Fatal("JWT_TTL_MINUTES is not a valid integer")
	}

	return Config{
		ServerPort: serverPort,
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     dbPort,
		DBUser:     getEnv("DB_USER", "user"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "sample"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
		JWTSecret:  getEnv("JWT_SECRET", ""),
		JWTTTL:     jwtTTL,
	}
}

func (c Config) DatabaseURL() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=UTC",
		c.DBHost,
		c.DBUser,
		c.DBPassword,
		c.DBName,
		c.DBPort,
		c.DBSSLMode,
	)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}
