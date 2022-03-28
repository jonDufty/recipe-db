package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/jonDufty/recipes/common/database"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	AuthEndpoint    string           `envconfig:"auth_endpoint"`
	GraphEndpoint   string           `envconfig:"graph_endpoint"`
	RecipesEndpoint string           `envconfig:"recipes_endpoint"`
	DB              *database.Config `envconfig:"database"`
}

type AuthConfig struct {
	Port int `envconfig:"port" default:"80"`
	AppConfig
}

type CookbookConfig struct {
	Port int `envconfig:"port" default:"80"`
	AppConfig
}

func NewAuthConfig() *AuthConfig {
	config := &AuthConfig{}
	baseDir := os.Getenv("BASE_DIR")

	if baseDir == "" {
		baseDir = "."
	}

	if err := LoadConfig(baseDir, "recipes", config); err != nil {
		log.Fatal(err)
	}

	err := LoadConfig(baseDir, "auth", config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}

func NewCookbookConfig() *CookbookConfig {
	config := &CookbookConfig{}
	baseDir := os.Getenv("BASE_DIR")

	if baseDir == "" {
		baseDir = "."
	}

	if err := LoadConfig(baseDir, "recipes", config); err != nil {
		log.Fatal(err)
	}

	err := LoadConfig(baseDir, "cookbook", config)
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
		log.Println("No .env file found. Using environment instead")
		return nil
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
