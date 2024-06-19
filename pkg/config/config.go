package config

import (
	"os"
	"strconv"
	"time"
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
