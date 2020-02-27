/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package log

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// LogrusLogger adapts a logrus logger to implement the Logger interface.
type LogrusLogger struct {
	entry *logrus.Entry
}

// LogrusLogger constructor.
func NewLogrusLogger() *LogrusLogger {
	return &LogrusLogger{
		entry: logrus.NewEntry(logrus.New()),
	}
}

// LogusLogger constructor with custom level a default formatter.
func NewDefaultLogrusLogger(strLevel string) *LogrusLogger {
	logger := NewLogrusLogger()
	if strLevel != "" {
		logrusLevel, _ := logrus.ParseLevel(strLevel)
		logger.GetLogger().SetLevel(logrusLevel)
	}
	logger.GetLogger().SetFormatter(&logrus.TextFormatter{
		FullTimestamp:          true,
		TimestampFormat:        "2006-01-02|15:04:05",
		DisableLevelTruncation: true,
	})
	return logger
}

// SetEntry allows to override the underlying logrus Entry.
func (l *LogrusLogger) SetEntry(entry *logrus.Entry) {
	l.entry = entry
}

// GetLogger returns the underlying logrus Logger
func (l LogrusLogger) GetLogger() *logrus.Logger {
	return l.entry.Logger
}

// Debug prints a message if the logger's level is at least debug.
func (l LogrusLogger) Debug(message string, keyValues ...interface{}) {
	l.withFields(keyValues...).Debug(message)
}

// Info prints a message if the logger's level is at least info.
func (l LogrusLogger) Info(message string, keyValues ...interface{}) {
	l.withFields(keyValues...).Info(message)
}

// Warn prints a message if the logger's level is at least warn.
func (l LogrusLogger) Warn(message string, keyValues ...interface{}) {
	l.withFields(keyValues...).Warn(message)
}

// Error prints a message if the logger's level is at least error.
func (l LogrusLogger) Error(message string, keyValues ...interface{}) {
	l.withFields(keyValues...).Error(message)
}

// With returns a new logger with the specified persistent keyValues.
func (l LogrusLogger) With(keyValues ...interface{}) Logger {
	logger := NewLogrusLogger()
	logger.SetEntry(l.withFields(keyValues...))
	return logger
}

// withFields converts variadic key-values into a logrus Fields.
func (l LogrusLogger) withFields(keyValues ...interface{}) *logrus.Entry {
	if len(keyValues)%2 != 0 {
		keyValues = append(keyValues, "MISSING_VALUE")
	}
	fields := make(map[string]interface{}, len(keyValues)/2)
	for i := 0; i < len(keyValues); i += 2 {
		k, v := keyValues[i], keyValues[i+1]
		fields[fmt.Sprint(k)] = v
	}
	return l.entry.WithFields(fields)
}
