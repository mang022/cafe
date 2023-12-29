package conf

import (
	"github.com/spf13/viper"
)

type config struct {
	Host struct {
		Address string
		Port    int
	}
	DB struct {
		Host string
		Port int
		User string
		Pwd  string
	}
	JWT struct {
		Secret string
	}
}

var Conf config

func SetupConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		panic(err)
	}
}
