package config

import (
	"github.com/spf13/viper"
)

type OptionsConfig struct {
	Component struct {
		Dir   string
		Props bool
	}
}

func GetConfig() OptionsConfig {
	var oc OptionsConfig

	viper.AddConfigPath(".")
	viper.SetConfigName("go-react")

	if err := viper.ReadInConfig(); err != nil {
		return oc
	}

	viper.Unmarshal(&oc)

	return oc
}
