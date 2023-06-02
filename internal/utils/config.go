package utils

import "github.com/spf13/viper"

type DbConfig struct {
	UrlDb string `mapstructure:"url_db"`
}

type Config struct {
	Db DbConfig `mapstructure:"db"`
}

var vp *viper.Viper

func LoadConfig() (Config, error) {
	vp = viper.New()
	var config Config

	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath("./internal/utils")
	err := vp.ReadInConfig()
	if err != nil {

		return Config{}, err
	}
	err = vp.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
