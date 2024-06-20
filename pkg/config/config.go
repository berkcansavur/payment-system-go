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
	GetIyzicoApiKey() string
	GetIyzicoApiSecret() string
	GetRandomString() string
	GetIyzicoConfig() IyzicoConfig
}

func GetApiURL() string {
	return os.Getenv("IYZICO_BASE_URL")
}

func GetIyzicoApiKey() string {
	return os.Getenv("IYZICO_API_KEY")
}

func GetIyzicoApiSecret() string {
	return os.Getenv("IYZICO_SECRET")
}
func GetMongoUri() string {
	return os.Getenv("MONGODB_URI")
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
		APIKey:    GetIyzicoApiKey(),
		APISecret: GetIyzicoApiSecret(),
		Rnd:       GetRandomString(),
	}
}
type Config struct {
	IyzicoApiKey    string
	IyzicoApiSecret string
	DbURI     		string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	log.Println("Configuration loaded successfully.")
	return &Config{
		IyzicoApiKey:    GetIyzicoApiKey(),
		IyzicoApiSecret: GetIyzicoApiSecret(),
		DbURI:     		 GetMongoUri(),
	}
}