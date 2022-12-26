package config

import (
	"log"

	"github.com/spf13/viper"
)

type OptionsConfig struct {
	Component struct {
		Dir string
		Props bool
	}
}

func GetConfig() OptionsConfig {
	var MkOptionsConfig OptionsConfig

	viper.AddConfigPath(".")
	viper.SetConfigName("make-react-file")
	
	if err := viper.ReadInConfig(); err != nil {
		log.Println("No config file found")
	}

	viper.Unmarshal(&MkOptionsConfig)

	return MkOptionsConfig
}