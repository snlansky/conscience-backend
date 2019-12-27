package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
)

type Server struct {
	Addr string `mapstructure:"addr"`
}

type UserCenter struct {
	PrivateUrl string `mapstructure:"private-url"`
	PublicUrl  string `mapstructure:"public-url"`
}

type Mongo struct {
	Uri string `mapstructure:"uri"`
	DB  string `mapstructure:"db"`
}

type Dao struct {
	Mongo Mongo `mapstructure:"mongo"`
}

type Config struct {
	Server     Server     `mapstructure:"server"`
	UserCenter UserCenter `mapstructure:"user-center"`
	Dao        Dao        `mapstructure:"dao"`
}

var GlobalConfig *Config

func Init() {
	var c Config
	err := viper.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	buf, err := json.MarshalIndent(&c, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))

	GlobalConfig = &c
}
