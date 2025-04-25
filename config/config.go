package config

import (
	"log"
	"os"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	RedisHost     string
	RedisPort     string
	RedisPassword string

	JWTSecret string
	RateLimit string
}

func LoadConfig() Config {

	log.Printf("DB_HOST: %s", os.Getenv("DB_HOST"))
	log.Printf("DB_PORT: %s", os.Getenv("DB_PORT"))
	log.Printf("DB_USER: %s", os.Getenv("DB_USER"))
	log.Printf("DB_PASSWORD: %s", os.Getenv("DB_PASSWORD"))
	log.Printf("DB_NAME: %s", os.Getenv("DB_NAME"))

	return Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),

		RedisHost:     os.Getenv("REDIS_HOST"),
		RedisPort:     os.Getenv("REDIS_PORT"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),

		JWTSecret: os.Getenv("JWT_SECRET"),
		RateLimit: os.Getenv("RATE_LIMIT"),
	}

}
