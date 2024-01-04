package cmd

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"mower-project/internal/interface/service"
)

const (
	configPath = "config/application.yml"
)

// Serve main application execution
func Serve() {
	config := initConfig()
	result := service.Execute(config)
	if result == nil {
		log.Error("file is empty")
		return
	}
	for _, mower := range result {
		log.Debug(map[string]interface{}{"X": mower.X, "Y": mower.Y, "Orientation": mower.Orientation})
		fmt.Printf("%d %d %s\n", mower.X, mower.Y, mower.Orientation)
	}
	log.Info("Program ended")
}

// initConfig get config values from application.yml file
func initConfig() *viper.Viper {
	log.Info("initConfig")
	var config = viper.New()
	config.AutomaticEnv()
	config.SetConfigFile(configPath)
	if err := config.ReadInConfig(); err != nil {
		log.Error("Cannot read config file")
		panic(err)
	}
	log.Info("Log file properly imported")
	return config
}
