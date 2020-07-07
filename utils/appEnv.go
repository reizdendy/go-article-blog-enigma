package utils

import (
	"log"

	"github.com/spf13/viper"
)

//GetCustomConf is a function of viper to get the custom env variables from file .env
func GetCustomConf(key, defVal string) string {
	viper.SetConfigFile("config.env")
	viper.AddConfigPath("")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("config file not found")
		log.Println("Default config has been loaded")
		return defVal
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}
