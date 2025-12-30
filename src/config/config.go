package config

import (
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Port	string `mapstructure:"PORT"`
	AppName		string	`mapstructure:"APP_NAME"`
	
	DB DatabaseConfig `mapstructure:db`
}

type DatabaseConfig struct {
	Dialect  string `mapstructure:"dialect"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"` 
	Charset  string `mapstructure:"charset"`
}

func LoadConfig() (config Config, err error) {
	if err := godotenv.Load(); err != nil {
		println("No .env file found (using system env vars or yaml defaults)")
	}
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return config, err
		}
	}

	viper.SetEnvPrefix("API")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	return
}

