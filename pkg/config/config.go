package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type ConfigurationProvider interface {
	GetApiURL() string
	GetApiKey() string
	GetApiSecret() string
	GetRandomString() string
	GetIyzicoConfig() IyzicoConfig
}

func GetApiURL() string {
	return os.Getenv("IYZICO_BASE_URL")
}

func GetApiKey() string {
	return os.Getenv("IYZICO_API_KEY")
}

func GetApiSecret() string {
	return os.Getenv("IYZICO_SECRET")
}

func GetRandomString() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

type IyzicoConfig struct {
	BaseURL   string
	APIKey    string
	APISecret string
	Rnd       string
}

func GetIyzicoConfig() *IyzicoConfig {
	return &IyzicoConfig{
		BaseURL:   GetApiURL(),
		APIKey:    GetApiKey(),
		APISecret: GetApiSecret(),
		Rnd:       GetRandomString(),
	}
}
type Config struct {
	ApiKey    string
	ApiSecret string
	DbURI     string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return &Config{
		ApiKey:    os.Getenv("IYZICO_API_KEY"),
		ApiSecret: os.Getenv("IYZICO_SECRET"),
		DbURI:     os.Getenv("MONGODB_URI"),
	}
}