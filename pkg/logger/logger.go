/*
 * File: /pkg/logger/logger.go                                                 *
 * Project: blog-service                                                       *
 * Created At: Tuesday, 2022/06/7 , 12:44:01                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/06/7 , 14:24:27                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

type Level int8

type Fields map[string]any

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
	PANIC
)

func (l Level) String() string {
	switch l {
	case DEBUG:
		return "debug"
	case INFO:
		return "info"
	case WARN:
		return "warn"
	case ERROR:
		return "error"
	case FATAL:
		return "fatal"
	case PANIC:
		return "panic"
	}

	return ""
}

type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	fields    Fields
	callers   []string
}

func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	return &Logger{newLogger: log.New(w, prefix, flag)}
}

func (l *Logger) clone() *Logger {
	nl := *l
	return &nl
}

func (l *Logger) WithLevel(level Level, v ...any) {
	l.Output(level, fmt.Sprint(v...))
}

func (l *Logger) WithFields(f Fields) *Logger {
	ll := l.clone()
	if ll.fields == nil {
		ll.fields = make(Fields)
	}

	for k, v := range f {
		ll.fields[k] = v
	}

	return ll
}

func (l *Logger) WithContext(ctx context.Context) *Logger {
	ll := l.clone()
	ll.ctx = ctx

	return ll
}

func (l *Logger) WithCaller(skip int) *Logger {
	ll := l.clone()
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		ll.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}

	return ll
}

func (l *Logger) WithCallerFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	callers := []string{}
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		s := fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function)

		callers = append(callers, s)
	}

	ll := l.clone()
	ll.callers = callers

	return ll
}

func (l *Logger) JSONFormat(level Level, message string) map[string]any {
	data := make(Fields, len(l.fields)+4)
	data["level"] = level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["callers"] = l.callers

	if len(l.fields) > 0 {
		for k, v := range l.fields {
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}
	}

	return data
}

func (l *Logger) Output(level Level, message string) {
	body, _ := json.Marshal(l.JSONFormat(level, message))
	content := string(body)

	switch level {
	case INFO, DEBUG, WARN, ERROR:
		l.newLogger.Print(content)
	case FATAL:
		l.newLogger.Fatal(content)
	case PANIC:
		l.newLogger.Panic(content)
	}
}

func (l *Logger) Info(v ...any) {
	l.Output(INFO, fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...any) {
	l.Output(INFO, fmt.Sprintf(format, v...))
}

func (l *Logger) Debug(v ...any) {
	l.Output(DEBUG, fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...any) {
	l.Output(DEBUG, fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(v ...any) {
	l.Output(WARN, fmt.Sprint(v...))
}

func (l *Logger) Warnf(format string, v ...any) {
	l.Output(WARN, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v ...any) {
	l.Output(ERROR, fmt.Sprint(v...))
}

func (l *Logger) Errorf(format string, v ...any) {
	l.Output(ERROR, fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(v ...any) {
	l.Output(FATAL, fmt.Sprint(v...))
}
func (l *Logger) Fatalf(format string, v ...any) {
	l.Output(FATAL, fmt.Sprintf(format, v...))
}

func (l *Logger) Panic(v ...any) {
	l.Output(PANIC, fmt.Sprint(v...))
}

func (l *Logger) Panicf(format string, v ...any) {
	l.Output(PANIC, fmt.Sprintf(format, v...))
}
