package server

import (
	"net/http"
	"time"
	"vk-chat-bot/src/utils"

	"github.com/sirupsen/logrus"
)

func Init() error {
	logrus.WithFields(logrus.Fields{
		"time": time.Now().Format(utils.TimeFormat),
	}).Info("Server started!")
	return http.ListenAndServe(":7272", nil)
}
