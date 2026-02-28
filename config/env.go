package config

import (
	"fmt"
	"os"
	"strconv"
)

type Env struct {
	DatabaseName string
	Password     string
	User         string
	Address      string
	Port         int
	APIPort      int
	MongoURI     string
}

func NewEnv() *Env {
	return &Env{
		DatabaseName: getStringEnv("MONGO_DATABASE", "barcode-api"),
		User:         getStringEnv("MONGO_USER", "root"),
		Password:     getStringEnv("MONGO_PASSWORD", "example"),
		Address:      getStringEnv("MONGO_ADDRESS", "localhost"),
		Port:         getIntEnv("MONGO_PORT", 27017),
		APIPort:      getIntEnv("API_PORT", 1111),
		MongoURI:     getStringEnv("MONGO_URI", ""),
	}
}

func (e *Env) ConnectionString() string {
	if e.MongoURI != "" {
		return e.MongoURI
	}

	return fmt.Sprintf("mongodb://%s:%s@%s:%d", e.User, e.Password, e.Address, e.Port)
}

func getStringEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getIntEnv(key string, fallback int) int {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return parsedValue
}
