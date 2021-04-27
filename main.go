package main

import (
	"net/http"
	"time"
	server "vk-chat-bot/src"
	"vk-chat-bot/src/modules/callbacks"
	"vk-chat-bot/src/modules/messages"
	"vk-chat-bot/src/modules/vk"
	"vk-chat-bot/src/utils"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	loadConfig()

	// VK API initialization
	vkapi := vk.Init()
	logrus.WithFields(logrus.Fields{
		"time": time.Now().Format(utils.TimeFormat),
	}).Info("VK started!")

	// VK callbacks initialization
	cb := callbacks.Init()
	logrus.WithFields(logrus.Fields{
		"time": time.Now().Format(utils.TimeFormat),
	}).Info("Callback server started!")

	// VK messages hanlders initialization
	if err := messages.Init(vkapi); err != nil {
		logrus.WithFields(logrus.Fields{
			"time":  time.Now().Format(utils.TimeFormat),
			"error": err.Error(),
		}).Error("Failed to init!")
	} else {
		logrus.WithFields(logrus.Fields{
			"time": time.Now().Format(utils.TimeFormat),
		}).Info("Messages handlers module started!")
	}

	// Server making
	http.HandleFunc("/callback", cb.HandleFunc)
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
