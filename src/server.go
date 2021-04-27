package server

import (
	"net/http"
	"time"
	"vk-chat-bot/src/utils"

	"github.com/sirupsen/logrus"
)

func Init(port *string) error {
	logrus.WithFields(logrus.Fields{
		"time": time.Now().Format(utils.TimeFormat),
	}).Info(
		"Server started on " + *port + "!",
	)
	return http.ListenAndServe(":"+*port, nil)
}
