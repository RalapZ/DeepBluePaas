package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func ParserConfig() {
	viper.AddConfigPath("./etc")
	viper.SetConfigName("config")
	viper.SetConfigFile("yaml")
	viper.ReadInConfig
	fmt.Printf("%##v", viper.GetString())
}
