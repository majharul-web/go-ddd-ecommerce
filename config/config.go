package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations *Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	JWTSecret   string
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("VERSION not set in .env file")
		os.Exit(1)
	}
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("SERVICE_NAME not set in .env file")
		os.Exit(1)
	}

	httpPortStr := os.Getenv("HTTP_PORT")
	httpPort, err := strconv.Atoi(httpPortStr)
	if err != nil {
		fmt.Println("Error parsing HTTP_PORT", err)
		os.Exit(1)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		fmt.Println("JWT_SECRET not set in .env file")
		os.Exit(1)
	}

	configurations = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    httpPort,
		JWTSecret:   jwtSecret,
	}

}

func GetConfig() *Config {
	if configurations == nil {
		loadConfig()
	}
	return configurations
}
