/*
 * File: /pkg/logger/logger.go                                                 *
 * Project: blog-service                                                       *
 * Created At: Tuesday, 2022/06/7 , 12:44:01                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/8 , 08:05:10                              *
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

	"github.com/gin-gonic/gin"
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
	level     Level
}

func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	return &Logger{newLogger: log.New(w, prefix, flag)}
}

func (l *Logger) clone() *Logger {
	nl := *l
	return &nl
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

func (l *Logger) WithLevel(level Level) *Logger {
	ll := l.clone()
	ll.level = level
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

func (l *Logger) WithTrace() *Logger {
	ginCtx, ok := l.ctx.(*gin.Context)

	if ok {
		return l.WithFields(Fields{
			"trace_id": ginCtx.MustGet("X-Trace-ID"),
			"span_id":  ginCtx.MustGet("X-Span-ID"),
		})
	}

	return l
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

func (l *Logger) Output(message string) {
	body, _ := json.Marshal(l.JSONFormat(l.level, message))
	content := string(body)

	switch l.level {
	case INFO, DEBUG, WARN, ERROR:
		l.newLogger.Print(content)
	case FATAL:
		l.newLogger.Fatal(content)
	case PANIC:
		l.newLogger.Panic(content)
	}
}

func (l *Logger) Info(ctx context.Context, v ...any) {
	l.WithLevel(INFO).WithContext(ctx).WithTrace()
	l.Output(fmt.Sprint(v...))
}

func (l *Logger) Infof(ctx context.Context, format string, v ...any) {
	l.WithLevel(INFO).WithContext(ctx).WithTrace()
	l.Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Debug(ctx context.Context, v ...any) {
	l.WithLevel(DEBUG).WithContext(ctx).WithTrace()
	l.Output(fmt.Sprint(v...))
}

func (l *Logger) Debugf(ctx context.Context, format string, v ...any) {
	l.WithLevel(DEBUG).WithContext(ctx).WithTrace()
	l.Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(ctx context.Context, v ...any) {
	l.WithLevel(WARN).WithContext(ctx).WithTrace()
	l.Output(fmt.Sprint(v...))
}

func (l *Logger) Warnf(ctx context.Context, format string, v ...any) {
	l.WithLevel(WARN).WithContext(ctx).WithTrace()
	l.Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Error(ctx context.Context, v ...any) {
	l.WithLevel(ERROR).WithContext(ctx).WithTrace()
	l.Output(fmt.Sprint(v...))
}

func (l *Logger) Errorf(ctx context.Context, format string, v ...any) {
	l.WithLevel(ERROR).WithContext(ctx).WithTrace()
	l.Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(ctx context.Context, v ...any) {
	l.WithLevel(FATAL).WithContext(ctx).WithTrace()
	l.Output(fmt.Sprint(v...))
}
func (l *Logger) Fatalf(ctx context.Context, format string, v ...any) {
	l.WithLevel(FATAL).WithContext(ctx).WithTrace()
	l.Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Panic(ctx context.Context, v ...any) {
	l.WithLevel(PANIC).WithContext(ctx).WithTrace()
	l.Output(fmt.Sprint(v...))
}

func (l *Logger) Panicf(ctx context.Context, format string, v ...any) {
	l.WithLevel(PANIC).WithContext(ctx).WithTrace()
	l.Output(fmt.Sprintf(format, v...))
}
