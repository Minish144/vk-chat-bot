package main

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.WithFields(logrus.Fields{
		"time": time.Now(),
	}).Info("Server started!")
	if err := http.ListenAndServe("7272", nil); err != nil {
		logrus.Fatal("Server is down!")
	}
}
