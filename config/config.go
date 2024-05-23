package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Postgres PostgresConfig
	Server   ServerConfig
}

type PostgresConfig struct {
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
}

type ServerConfig struct {
	Host string
	Port string
}

func Load(path string) Config {
	viper.SetConfigFile(".env")

	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	return Config{
		Postgres: PostgresConfig{
			DbHost:     viper.Get("DB_HOST").(string),
			DbPort:     viper.Get("DB_PORT").(string),
			DbName:     viper.Get("DB_NAME").(string),
			DbUser:     viper.Get("DB_USER").(string),
			DbPassword: viper.Get("DB_PASSWORD").(string),
		},
		Server: ServerConfig{
			Host: viper.Get("SERVER_HOST").(string),
			Port: viper.Get("SERVER_PORT").(string),
		},
	}
}
