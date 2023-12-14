package config

import (
	"github.com/spf13/viper"
	"sync"
)

type Config struct {
	DBDriver     string `mapstructure:"DB_DRIVER"`
	DBUrl        string `mapstructure:"DB_URL"`
	Port         string `mapstructure:"PORT"`
	OpenAIAPIKey string `mapstructure:"OPENAI_API_KEY"`
}

var once sync.Once
var config Config

func GetConfig() Config {
	once.Do(func() {
		var err error
		config, err = loadConfig(".")
		if err != nil {
			panic(err)
		}
	})

	return config
}

func loadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	return config, err
}
