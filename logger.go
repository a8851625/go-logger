package logger

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	Debug   *Logger
	Info    *Logger
	Warning *Logger
	Error   *Logger
	once    sync.Once
)

// Logger wraps logrus Logger
type Logger struct {
	logger *logrus.Logger
	level  logrus.Level
}

func init() {
	once.Do(func() {
		Debug = newLogger(logrus.DebugLevel)
		Info = newLogger(logrus.InfoLevel)
		Warning = newLogger(logrus.WarnLevel)
		Error = newLogger(logrus.ErrorLevel)
	})
}

func newLogger(level logrus.Level) *Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(level)

	return &Logger{
		logger: log,
		level:  level,
	}
}

// Println logs a message at the logger's level.
func (l *Logger) Println(v ...interface{}) {
	switch l.level {
	case logrus.DebugLevel:
		l.logger.Debug(v...)
	case logrus.InfoLevel:
		l.logger.Info(v...)
	case logrus.WarnLevel:
		l.logger.Warn(v...)
	case logrus.ErrorLevel:
		l.logger.Error(v...)
	}
}

// Printf logs a formatted message at the logger's level.
func (l *Logger) Printf(format string, v ...interface{}) {
	switch l.level {
	case logrus.DebugLevel:
		l.logger.Debugf(format, v...)
	case logrus.InfoLevel:
		l.logger.Infof(format, v...)
	case logrus.WarnLevel:
		l.logger.Warnf(format, v...)
	case logrus.ErrorLevel:
		l.logger.Errorf(format, v...)
	}
}

// Fatalln logs a message at level Fatal then the process will exit with status set to 1.
func (l *Logger) Fatalln(v ...interface{}) {
	l.logger.Fatalln(v...)
}

// Fatalf logs a formatted message at level Fatal then the process will exit with status set to 1.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.logger.Fatalf(format, v...)
}

// GetLevel returns the current logger level.
func (l *Logger) GetLevel() logrus.Level {
	return l.level
}

// SetOutput sets the logger output.
func (l *Logger) SetOutput(output *os.File) {
	l.logger.SetOutput(output)
}
