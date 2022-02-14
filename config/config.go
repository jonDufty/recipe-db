package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	AuthEndpoint    string `envconfig:"auth_endpoint"`
	GraphEndpoint   string `envconfig:"graph_endpoint"`
	RecipesEndpoint string `envconfig:"recipes_endpoint"`
}

type AuthConfig struct {
	Port int
	AppConfig
}

func NewAuthConfig() *AuthConfig {
	config := &AuthConfig{}
	baseDir := os.Getenv("BASE_DIR")
	err := LoadConfig(baseDir, "auth", config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

func LoadConfig(baseDir string, envPrefix string, config interface{}) error {
	configFile := ".env"
	configPath := filepath.Join(baseDir, configFile)
	fmt.Printf("Loading .env file %s\n", configPath)

	if _, err := os.Stat(configPath); err != nil {
		return errors.New("Error loading config file." + err.Error())
	}

	err := godotenv.Load(configPath)
	if err != nil {
		return err
	}

	err = envconfig.Process(envPrefix, config)
	if err != nil {
		return err
	}

	return nil
}
