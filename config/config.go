package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	HTTP struct {
		Port string `yaml:"port"`
	} `yaml:"http"`

	Mongodb struct {
		URL        string `yaml:"url"`
		Database   string `yaml:"database"`
		Collection string `yaml:"collection"`
	} `yaml:"mongodb"`
}

func NewConfig() Config {
	cfg := Config{}
	file, err := os.ReadFile("./config/config.yml")
	if err != nil {
		log.Fatalln(err)
	}
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		log.Fatalln(err)
	}
	return cfg
}
