package logs

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"sactl/config"
)

// 记录日志
func Logs() {
	optins, err := config.ParseConfig()
	if err != nil {
		ErrInfo("Logs optins", "fail", err)
	}

	logger := lumberjack.Logger{
		Filename: optins.Log.FileName,
		MaxSize:  optins.Log.Max_size,
		MaxAge:   optins.Log.Max_age,
		Compress: optins.Log.Compress,
	}

	defer logger.Close()
	logrus.SetOutput(&logger)

	level, err := logrus.ParseLevel(optins.Log.Level)
	if err != nil {
		ErrInfo("log Level", "fail", err)
	}
	logrus.SetLevel(logrus.Level(level))
	logrus.SetReportCaller(true)
	//logrus.SetFormatter(&logrus.JSONFormatter{})
}

func ErrInfo(action string, sample interface{}, err error) {
	if err != nil {
		logrus.WithFields(
			logrus.Fields{
				"Action":     action,
				"Sample":     sample,
				"Error_Info": err,
			}).Error("SaCtl")
	} else {
		logrus.WithFields(
			logrus.Fields{
				"action": action,
				"sample": sample,
				"Info":   err,
			}).Info("SaCtl")
	}
}
