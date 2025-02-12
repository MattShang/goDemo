package main

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

var Conf Config

func init() {
	// 读取配置文件
	yamlFile, err := os.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	// 解析配置文件
	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		panic(err)
	}
}
