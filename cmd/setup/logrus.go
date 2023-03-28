package setup

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func ConfigureLogrus() {
	env := os.Getenv("ENV")
	lum := &lumberjack.Logger{
		Filename:   "log/server-" + env + ".log",
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	}

	logrus.SetOutput(lum)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceQuote:    true,
	})
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Info("logrus & lumberjack activated")
}
