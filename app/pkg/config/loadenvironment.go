package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var Config map[string]interface{}

const (
	CONFIG_ENV  = "env-dev"
	CONFIG_YAML = "config-dev"
)

func Loadenvironment() error {
	confpath := "../../config/."

	if err := godotenv.Load(confpath + CONFIG_ENV); err != nil {
		log.Printf("Error loading .env file: %s", err.Error())
		return err
	}

	v := viper.New()
	v.SetConfigName(CONFIG_YAML)
	v.SetConfigType("yaml")
	v.AddConfigPath(confpath)

	if err := v.ReadInConfig(); err != nil {
		log.Printf("Error reading YAML config file: %s", err.Error())
		return err
	}

	Config = v.AllSettings()
	return nil
}
