package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

const (
	ConfigFilePath = "./internal/config/"
	ConfigFileName = "config"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server" validate:"required"`
	Postgres PostgresConfig `mapstructure:"postgres" validate:"required"`
}

type ServerConfig struct {
	Host string `mapstructure:"host" validate:"required"`
	Port string `mapstructure:"port" validate:"required"`
}

type PostgresConfig struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     string `mapstructure:"port" validate:"required"`
	User     string `mapstructure:"user" validate:"required"`
	DBName   string `mapstructure:"dbname" validate:"required"`
	SSLMode  string `mapstructure:"sslmode" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
}

func NewConfig(configFilePath string, configFileName string) (Config, error) {
	myViper := viper.New()
	myViper.AddConfigPath(configFilePath)
	myViper.SetConfigName(configFileName)

	if err := myViper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var result Config
	err := myViper.Unmarshal(&result)
	if err != nil {
		return Config{}, err
	}
	validate := validator.New()
	if err := validate.Struct(&result); err != nil {
		return Config{}, err
	}
	return result, nil
}
