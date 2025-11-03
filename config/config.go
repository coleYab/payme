package config

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	MONGO_DATABASE string
	MONGO_URI string
	REDIS_ADDR string
	REDIS_PASS string
	REDIS_DB string
	ADDR string
}

func reqEnv(name string) string {
	value, ok := os.LookupEnv(name)
	if !ok || len(value) == 0 {
		log.Fatalln(name + " is not set");
	}

	return value
}

func getenvFallback(name, fallback string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		return fallback
	}

	return value
}

var CFG = Config{
	MONGO_URI: reqEnv("MONGO_URI"),
	MONGO_DATABASE: getenvFallback("MONGO_DATABASE", "payme"),
	REDIS_PASS: getenvFallback("REDIS_PASS", ""),
	REDIS_ADDR: getenvFallback("REDIS_ADDR", "localhost:6379"),
	ADDR: getenvFallback("ADDR", ":8080"),
};
