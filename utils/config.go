package utils

import "github.com/spf13/viper"

type Config struct {
	MODE             string `mapstructure:"MODE"`
	ADDR             string `mapstructure:"ADDR"`
	SECRET_KEY_TOKEN string `mapstructure:"SECRET_KEY_TOKEN"`
	HOST             string `mapstructure:"HOST"`
	PORT             string `mapstructure:"PORT"`
	PASSWD           string `mapstructure:"PASSWD"`
	EMAIL            string `mapstructure:"EMAIL"`
	EXPIRES_AT       string `mapstructure:"EXPIRES_AT"`

	POSTGRES_HOST     string `mapstructure:"POSTGRES_HOST"`
	POSTGRES_USERNAME string `mapstructure:"POSTGRES_USERNAME"`
	POSTGRES_PASSWORD string `mapstructure:"POSTGRES_PASSWORD"`
	POSTGRES_DBNAME   string `mapstructure:"POSTGRES_DBNAME"`
}

var AppConfig *Config

func init() {
	viper.AutomaticEnv()

	AppConfig = &Config{
		MODE:              viper.GetString("MODE"),
		ADDR:              viper.GetString("ADDR"),
		SECRET_KEY_TOKEN:  viper.GetString("SECRET_KEY_TOKEN"),
		HOST:              viper.GetString("HOST"),
		PORT:              viper.GetString("PORT"),
		PASSWD:            viper.GetString("PASSWD"),
		EMAIL:             viper.GetString("EMAIL"),
		EXPIRES_AT:        viper.GetString("EXPIRES_AT"),
		POSTGRES_HOST:     viper.GetString("POSTGRES_HOST"),
		POSTGRES_USERNAME: viper.GetString("POSTGRES_USERNAME"),
		POSTGRES_PASSWORD: viper.GetString("POSTGRES_PASSWORD"),
		POSTGRES_DBNAME:   viper.GetString("POSTGRES_DBNAME"),
	}
}
