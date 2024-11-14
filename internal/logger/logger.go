package logger

import (
    "os"
    "time"

    "github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
    log = logrus.New()
    log.SetFormatter(&logrus.JSONFormatter{
        TimestampFormat: time.RFC3339,
    })
    log.SetOutput(os.Stdout)
}

// GetLogger returns the logger instance
func GetLogger() *logrus.Logger {
    return log
}

// Info logs info level messages
func Info(msg string, fields map[string]interface{}) {
    if fields != nil {
        log.WithFields(fields).Info(msg)
    } else {
        log.Info(msg)
    }
}

// Error logs error level messages
func Error(msg string, err error, fields map[string]interface{}) {
    if fields == nil {
        fields = make(map[string]interface{})
    }
    fields["error"] = err
    log.WithFields(fields).Error(msg)
}