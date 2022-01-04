package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)


var Config *viper.Viper

func ParserConfig() {
	fmt.Println(os.Getwd())
	Config=viper.New()
	Config.AddConfigPath("/Users/ralap/code/golang/deepblue/DeepBluePaas/etc")
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	if err := Config.ReadInConfig(); err!=nil {
		panic(err)
	}
	//fmt.Printf("%##v",config)
	//fmt.Printf("%##v", Config.Get("db.datasource"))
}
