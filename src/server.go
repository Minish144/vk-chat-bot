package server

import (
	"net/http"
	"time"
	"vk-chat-bot/src/utils"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Init() error {
	serverPort := viper.GetString("vk.port")
	logrus.WithFields(logrus.Fields{
		"time": time.Now().Format(utils.TimeFormat),
	}).Info(
		"Server started on " + serverPort + "!",
	)
	return http.ListenAndServe(":"+serverPort, nil)
}
