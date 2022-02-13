package common

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Configuration exported
type Configuration struct {
	Server   ServerConfiguration   `mapstructure:"server"`
	Database DatabaseConfiguration `mapstructure:"database"`
}

// ServerConfiguration exported
type ServerConfiguration struct {
	Port string `mapstructure:"port"`
}

// DatabaseConfiguration exported
type DatabaseConfiguration struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

// New initializes a new Configuration from the ENV variables
func LoadConfig() *Configuration {
	viper.SetConfigName("config")

	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	var config Configuration

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error while reading config file: %v", err))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Configuration file changed")
	})

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode config file to struct, err: %v", err)
	}

	return &config
}
