package logger

import (
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

const (
	LevelDebug = logrus.DebugLevel
	LevelInfo  = logrus.InfoLevel
)

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	Fatal(msg string, args ...any)
}

type logger struct {
	log *logrus.Logger
}

func NewLogger() *logger {
	logger := &logger{logrus.New()}
	return logger.WithFormatter(&prefixed.TextFormatter{
		SpacePadding: 10,
	})
}

func (l *logger) WithFormatter(formatter logrus.Formatter) *logger {
	l.log.SetFormatter(formatter)
	return l
}

func (l *logger) WithOutput(output io.Writer) *logger {
	l.log.SetOutput(output)
	return l
}

func (l *logger) WithLevel(level logrus.Level) *logger {
	l.log.SetLevel(level)
	return l
}

func (l *logger) Debug(msg string, args ...any) {
	l.log.WithFields(parseFields(args)).Debugln(msg)
}

func (l *logger) Info(msg string, args ...any) {
	l.log.WithFields(parseFields(args)).Infoln(msg)
}

func (l *logger) Warn(msg string, args ...any) {
	l.log.WithFields(parseFields(args)).Warnln(msg)
}

func (l *logger) Error(msg string, args ...any) {
	l.log.WithFields(parseFields(args)).Errorln(msg)
}

func (l *logger) Fatal(msg string, args ...any) {
	l.log.WithFields(parseFields(args)).Fatalln(msg)
}

func parseFields(args []any) logrus.Fields {
	fields := logrus.Fields{}

	var key string
	for i, arg := range args {
		if i%2 == 0 {
			key = fmt.Sprintf("\n\t%v", arg)
		} else {
			fields[key] = fmt.Sprintf("%v", arg)
		}
	}

	return fields
}
