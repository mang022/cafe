package main

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
}

var conf config

func setupConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}
}
