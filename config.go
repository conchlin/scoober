package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

// the specific database parameters for a connection
type dbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

// the overarching categories found within the config.toml
type Config struct {
	Database dbConfig
}

// looks for the config.toml file and loads the parameters into the various structs
func LoadConfig() *Config {
	tomlFile := &Config{}

	if _, err := toml.DecodeFile("config.toml", tomlFile); err != nil {
		log.Fatal(err)
	}

	return tomlFile
}
