package logs

import (
	"demo3/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

func Log() {
	options, err := config.ParseConfig()
	if err != nil {
		log.Panic(err)
	}

	logger := lumberjack.Logger{
		Filename: options.Log.FileName,
		MaxSize:  options.Log.Max_size,
		MaxAge:   options.Log.Max_age,
		Compress: options.Log.Compress,
	}

	defer logger.Close()
	logLevel, err := logrus.ParseLevel(options.Log.Level)

	logrus.SetOutput(&logger)

	// 必须通过 level 才能将其写入到文件中
	logrus.SetLevel(logLevel)
	//logrus.SetReportCaller(true)

}

func WithFields(metrics, sample interface{}) {
	logrus.WithFields(
		logrus.Fields{
			"指标": metrics,
			"数据": sample,
		}).Debug("查询结果")
}
