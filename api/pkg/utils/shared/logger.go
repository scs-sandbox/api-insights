// Copyright 2022 Cisco Systems, Inc. and its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package shared

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"runtime"
	"strings"
)

// LogOption defines the cli options for service
type LogOption struct {
	LogFormat string `long:"log-format"  description:"Log Format - text or json" default:"json" env:"LOG_FORMAT"`
	LogLevel  string `long:"log-level"   description:"Log Level - debug, info, warn, error, or fatal" default:"info" env:"LOG_LEVEL"`
	LogName   string `long:"log-name"    description:"Log Name - log name" env:"LOG_NAME"`
}

// Logger defines a generic logging interface
type Logger interface {
	Debugf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Print(v ...interface{})
	Println(v ...interface{})
	WithField(key string, value interface{}) Logger
	SetLevel(level string) error
	GetLevel() log.Level
}

// LogOpt is global logging configuration instance
var LogOpt = LogOption{LogFormat: "json", LogLevel: "info", LogName: ""}
var globalLogger Logger
var hostname, _ = os.Hostname()

func init() {
	if os.Getenv("LOG_LEVEL") != "" {
		LogOpt.LogLevel = os.Getenv("LOG_LEVEL")
	}
	if os.Getenv("LOG_FORMAT") != "" {
		LogOpt.LogFormat = os.Getenv("LOG_FORMAT")
	}

	if LogOpt.LogFormat == "json" {
		log.SetFormatter(&log.JSONFormatter{TimestampFormat: "2006-01-02T15:04:05"})
	} else {
		textFmt := log.TextFormatter{
			ForceColors:      false,
			DisableColors:    true,
			DisableTimestamp: false,
			FullTimestamp:    true,
			TimestampFormat:  "2006-01-02T15:04:05",
			DisableSorting:   false,
		}
		log.SetFormatter(&textFmt)
	}
	log.SetOutput(os.Stderr)
	lvl, err := log.ParseLevel(LogOpt.LogLevel)
	if err != nil {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(lvl)
	}

	globalLogger = NewLogger(log.Fields{})
}

// NewLogger will create a new Logger instance
func NewLogger(additionalFields log.Fields) Logger {
	return NewLoggerWithSkip(additionalFields, 4)
}

// NewLoggerWithSkip will create a new Logger instance
func NewLoggerWithSkip(additionalFields log.Fields, skip int) Logger {
	fields := log.Fields{
		"service": LogOpt.LogName,
		"host":    hostname,
	}
	for k, v := range additionalFields {
		fields[k] = v
	}
	entry := log.WithFields(fields)
	return logrusLogger{skip: skip, entry: entry}
}

// GlobalLogger will return a global logger
func GlobalLogger() Logger {
	return globalLogger
}

// LoggerFromRequest will return a Logger instance from request object
func LoggerFromRequest(r *http.Request) Logger {
	if r == nil {
		return globalLogger
	}
	traceID := r.Header.Get(string(APIInsightsTraceID))
	if traceID != "" {
		return NewLoggerWithSkip(log.Fields{"trace": traceID}, 3)
	}
	return LoggerFromContext(r.Context())
}

// LoggerFromContext will return a Logger instance from context object
func LoggerFromContext(ctx context.Context) Logger {
	if ctx == nil {
		return globalLogger
	}
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		spanCtx := span.Context()
		return NewLoggerWithSkip(log.Fields{"trace": fmt.Sprintf("%s", spanCtx)}, 3)
	}
	return globalLogger
}

// LogErrorf logs message in error level
func LogErrorf(format string, v ...interface{}) {
	globalLogger.Errorf(format, v...)
}

// LogInfof logs message in info level
func LogInfof(format string, v ...interface{}) {
	globalLogger.Infof(format, v...)
}

// LogDebugf logs message in debug level
func LogDebugf(format string, v ...interface{}) {
	globalLogger.Debugf(format, v...)
}

// LogFatalf logs message in fatal level
func LogFatalf(format string, v ...interface{}) {
	globalLogger.Fatalf(format, v...)
}

// LogWarnf logs message in warn level
func LogWarnf(format string, v ...interface{}) {
	globalLogger.Warnf(format, v...)
}

// LogPrintf logs message in error level
func LogPrintf(format string, v ...interface{}) {
	globalLogger.Printf(format, v...)
}

type logrusLogger struct {
	skip  int
	entry *log.Entry
}

func (l logrusLogger) entryWithCallInfo() *log.Entry {
	file, line, fn := l.callerInfo()
	return l.entry.WithFields(log.Fields{"file": file, "line": line, "func": fn})
}

func (l logrusLogger) Debugf(format string, v ...interface{}) {
	l.entryWithCallInfo().Debugf(format, v...)
}

func (l logrusLogger) Fatalf(format string, v ...interface{}) {
	l.entryWithCallInfo().Fatalf(format, v...)
}

func (l logrusLogger) Infof(format string, v ...interface{}) {
	l.entryWithCallInfo().Infof(format, v...)
}

func (l logrusLogger) Errorf(format string, v ...interface{}) {
	l.entryWithCallInfo().Errorf(format, v...)
}

func (l logrusLogger) Warnf(format string, v ...interface{}) {
	l.entryWithCallInfo().Warnf(format, v...)
}

func (l logrusLogger) Printf(format string, v ...interface{}) {
	l.entryWithCallInfo().Printf(format, v...)
}

func (l logrusLogger) Print(v ...interface{}) {
	l.entryWithCallInfo().Print(v...)
}

func (l logrusLogger) Println(v ...interface{}) {
	l.entryWithCallInfo().Println(v...)
}

func (l logrusLogger) SetLevel(lvl string) error {
	level, err := log.ParseLevel(lvl)
	if err != nil {
		return err
	}
	l.entry.Logger.SetLevel(level)
	return nil
}

func (l logrusLogger) GetLevel() log.Level {
	return l.entry.Logger.Level
}

func (l logrusLogger) WithField(key string, value interface{}) Logger {
	return logrusLogger{
		skip:  l.skip - 1,
		entry: l.entry.WithFields(log.Fields{key: value}),
	}
}

// Retrieve caller info, file name and line number
func (l logrusLogger) callerInfo() (string, int, string) {
	pc, file, line, _ := runtime.Caller(l.skip)
	index := strings.Index(file, "api-insights")
	if index != -1 {
		file = file[index:]
	}
	fn := runtime.FuncForPC(pc).Name()
	index = strings.Index(fn, "api-insights")
	if index != -1 {
		fn = fn[index:]
	}
	return file, line, fn
}
