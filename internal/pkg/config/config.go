//nolint:tagliatelle
package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	cfgPath = "./internal/pkg/config/config.yaml"
)

type ServerConfig struct {
	Host               string `yaml:"host"`
	Port               string `yaml:"port"`
	StorageType        string `yaml:"storage_type"`
	DBConnectingURL    string `yaml:"db_connection_url"`
	DBMigrationsFolder string `yaml:"db_migrations_folder"`
}

type Config struct {
	Server ServerConfig `yaml:"server"`
}

func ReadConfig() *Config {
	cfg := &Config{}

	file, err := os.Open(cfgPath)
	if err != nil {
		log.Println("Something went wrong while opening config file:", err)

		return nil
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		log.Println("Something went wrong while reading config file:", err)

		return nil
	}

	log.Println("Successfully opened config")

	return cfg
}
