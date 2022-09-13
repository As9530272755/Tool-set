package controller

import (
	"github.com/spf13/viper"
	"log"
)

var (
	User     string
	Passwrod string
)

func Config() (url string) {
	viper.SetConfigFile("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./etc")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		return ""
	}

	urlPath := viper.GetString("web.url")
	port := viper.GetString("web.port")
	User = viper.GetString("login.user")
	Passwrod = viper.GetString("login.password")
	return urlPath + port
}
