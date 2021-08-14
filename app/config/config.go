package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Debug  bool
	Server struct {
		Address string
	}
	Context struct {
		Timeout int
	}
	Database struct {
		Host string
		Port string
		User string
		Pass string
		Name string
	}
	JWT struct {
		Secret  string
		Expired int
	} `mapstructure:"jwt"`
}

func GetConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(os.Getenv("APP_PATH") + "/app/config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		panic(err)
	}
	return c
}
