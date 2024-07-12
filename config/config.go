package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbUserName string
	DBPassword string
	DBAddress  string
	Port       string
	DBName     string
	PublicHost string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load the DB envs")
		log.Fatal(err)
	}

}

func GetConfigs() *Config {
	return &Config{
		DbUserName: getEnv("DB_USER_NAME", "root"),
		DBPassword: getEnv("DB_PASS", "password"),
		DBAddress:  getEnv("DB_ADDRESS", "127.0.0.1:3306"),
		Port:       getEnv("PORT", ":8080"),
		DBName:     getEnv("DB_NAME", "User_DB"),
		PublicHost: getEnv("PUBLIC_HOST", "localhost"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
