package logger

import "github.com/sirupsen/logrus"

type Logger interface {
	Error(err error)
	Errorf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	WithFields(fields logrus.Fields) Logger
	WithField(key string, value interface{}) Logger
	SetReportCaller(reportCaller bool)
	Logurs() *logrus.Logger
}
