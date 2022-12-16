package config

import (
	"github.com/spf13/viper"
)

const DEFAULT_PATH = "./"

type Config struct {
	Server Server `yaml:"server"`
	Redis  Redis  `yaml:"redis"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
