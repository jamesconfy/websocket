package utils

import "github.com/spf13/viper"

type Config struct {
	ADDR             string `mapstructure:"ADDR"`
	SECRET_KEY_TOKEN string `mapstructure:"SECRET_KEY_TOKEN"`
	HOST             string `mapstructure:"HOST"`
	PORT             string `mapstructure:"PORT"`
	PASSWD           string `mapstructure:"PASSWD"`
	EMAIL            string `mapstructure:"EMAIL"`

	POSTGRES_HOST     string `mapstructure:"POSTGRES_HOST"`
	POSTGRES_USERNAME string `mapstructure:"POSTGRES_USERNAME"`
	POSTGRES_PASSWORD string `mapstructure:"POSTGRES_PASSWORD"`
	POSTGRES_DBNAME   string `mapstructure:"POSTGRES_DBNAME"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
