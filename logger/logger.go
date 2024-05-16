package logger

import (
	"context"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Error(ctx context.Context, err error)
	Errorf(ctx context.Context, format string, args ...interface{})
	Info(ctx context.Context, args ...interface{})
	Infof(ctx context.Context, format string, args ...interface{})
	Warn(ctx context.Context, args ...interface{})
	Warnf(ctx context.Context, format string, args ...interface{})
	Debug(ctx context.Context, args ...interface{})
	Debugf(ctx context.Context, format string, args ...interface{})
	WithFields(fields logrus.Fields) Logger
	WithField(key string, value interface{}) Logger
	// WithContext(ctx context.Context)Logger
	SetReportCaller(reportCaller bool)
	Logurs() *logrus.Logger
}


