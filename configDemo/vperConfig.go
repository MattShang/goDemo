package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	appName := viper.GetString("app.name")
	fmt.Printf("使用 viper 解析 yaml 文件: %s \n", appName)

	fmt.Println("使用 yaml.v3 包解析 yaml 文件: ", Conf.Server.Host)
}
