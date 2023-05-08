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

var AppConfig Config

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		panic(err)
	}
}
