package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBPort     string `mapstructure:"DB_PORT"`
	ACCOUNTSID string `mapstructure:"DB_ACCOUNTSID"`
	SERVICEID  string `mapstructure:"DB_SERVICEID"`
	AUTHTOKEN  string `mapstructure:"DB_AUTHTOKEN"`
	UNIDOCKEY  string `mapstructure:"UNIDOC_LICENSE_API_KEY"`
}

var envs = []string{
	"DB_HOST", "DB_NAME", "DB_USER", "DB_PASSWORD", "DB_PORT", "DB_ACCOUNTSID", "DB_SERVICEID", "DB_AUTHTOKEN", "DB_UNIDOC_LICENSE_API_KEY",
}

func LoadConfig() (Config, error) {
	var config Config
	
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.AutomaticEnv()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error while loading the env file..")
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
