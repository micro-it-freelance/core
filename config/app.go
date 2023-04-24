package core_config

import "github.com/spf13/viper"

var APP struct {
	Name string `mapstructure:"APP_NAME"`
}

func init() {
	APP.Name = viper.GetString("APP_NAME")
}