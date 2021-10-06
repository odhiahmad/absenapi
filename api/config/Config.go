package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
)

var (
	AppConfig *Config
)

type Config struct {
	CfgViper *viper.Viper
	Logger   logrus.Logger
}


func NewConfig() *Config {
	c := new(Config)
	c.init()
	AppConfig = c
	return c
}

func (c *Config) init() {
	c.CfgViper = viper.GetViper()
	c.CfgViper.AutomaticEnv()
	c.CfgViper.SetConfigFile(".env")
	err := c.CfgViper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	level := c.CfgViper.GetString("LOG_LEVEL")
	c.logger(level)
}

func (c *Config) logger(level string) {
	logFormat := new(logrus.JSONFormatter)
	var logLevel, err = logrus.ParseLevel(level)
	if err != nil {
		panic(err)
	}
	c.Logger = logrus.Logger{
		Out:       os.Stderr,
		Formatter: logFormat,
		Level:     logLevel,
	}
}

func (c *Config) GetConfigValue(key string,defaultValue string) string {
	if envVal := viper.GetString(key); len(envVal) != 0 {
		return envVal
	}
	return defaultValue
}
