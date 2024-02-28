package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	APIPORT string `mapstructure:"API_PORT"`
}

func Init() *Config {
	fmt.Println("Inside Config directory")
	viper.SetConfigFile("local")
	viper.AddConfigPath("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Panicln(err)
	}

	config := Config{}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Panicln("Error marshaling config file")
	}
	return &config
}
