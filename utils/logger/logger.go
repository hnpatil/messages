package logger

import "github.com/sirupsen/logrus"

func GetInstance() logrus.FieldLogger {
	lg := logrus.New()

	lg.SetFormatter(&logrus.TextFormatter{})

	return lg
}
