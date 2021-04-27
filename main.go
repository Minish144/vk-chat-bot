package main

import (
	"time"
	server "vk-chat-bot/src"

	"github.com/sirupsen/logrus"
)

func main() {
	if err := server.Init(); err != nil {
		logrus.WithFields(logrus.Fields{
			"time":  time.Now().Format("2-01-2006 15:04:05"),
			"error": err.Error(),
		}).Fatal("Server is down!")
	}
}
