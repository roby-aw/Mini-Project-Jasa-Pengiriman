package config

import (
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type AppConfig struct {
	App struct {
		Port int `toml:"port"`
	} `toml:"app"`
	Database struct {
		DBURL string `toml:"dburl"`
	} `toml:"database"`
	Secrettoken struct {
		Token string `toml:"token"`
	} `toml:"secretToken"`
	RajaOngkir struct {
		Api_key string `toml:"api_key"`
	} `toml:"RajaOngkir"`
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}
	return appConfig
}

func initConfig() *AppConfig {
	var defaultconfig AppConfig
	defaultconfig.App.Port = 8080
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Info("Error when loading config", err)
		return &defaultconfig
	}

	var finalConfig AppConfig
	err := viper.Unmarshal(&finalConfig)
	if err != nil {
		log.Info("Error when parse config", err)
	}
	return &finalConfig
}
