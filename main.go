package main

import (
	"fmt"
	"time"
	server "vk-chat-bot/src"
	"vk-chat-bot/src/modules/vk"
	"vk-chat-bot/src/utils"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	loadConfig()

	vkApi := vk.Init()
	fmt.Println(vkApi)
	logrus.WithFields(logrus.Fields{
		"time": time.Now().Format(utils.TimeFormat),
	}).Info("VK started!")

	if err := server.Init(); err != nil {
		logrus.WithFields(logrus.Fields{
			"time":  time.Now().Format(utils.TimeFormat),
			"error": err.Error(),
		}).Fatal("Server is down!")
	}
}

func loadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		logrus.WithFields(logrus.Fields{
			"name": e.Name,
		}).Info("Config file changed!")
	})

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logrus.Fatalln("Config file not found")
		} else {
			logrus.WithFields(logrus.Fields{
				"time":  time.Now().Format(utils.TimeFormat),
				"error": err,
			}).Fatalln("Error while reading config!")
		}
	}

	logrus.WithFields(logrus.Fields{
		"time": time.Now().Format(utils.TimeFormat),
	}).Infoln("Config loaded!")
}
