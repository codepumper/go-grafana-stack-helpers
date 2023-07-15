package logging

import (
	"fmt"
	"runtime"
	"time"
	"sync"

	"github.com/afiskon/promtail-client/promtail"
)

type Logger struct {
	loki promtail.Client
}

var (
	lock sync.Mutex
	loggerInstance *Logger
)

func InitLogger(format, labels string, config *promtail.ClientConfig) error {
	var err error
	loggerInstance, err = NewLogger(format, labels, config)
	return err
}

func NewLogger(format, labels string, config *promtail.ClientConfig) (*Logger, error) {
	if format != "proto" && format != "json" {
		return nil, fmt.Errorf("invalid log format: %s", format)
	}

	if(config == nil) {
		config = &promtail.ClientConfig{
			PushURL:            "http://localhost:3100/api/prom/push",
			Labels:             labels,
			BatchWait:          5 * time.Second,
			BatchEntriesNumber: 10000,
			SendLevel:          promtail.INFO,
			PrintLevel:         promtail.ERROR,
		}
	}

	var (
		loki promtail.Client
		err  error
	)

	if format == "proto" {
		loki, err = promtail.NewClientProto(*config)
	} else {
		loki, err = promtail.NewClientJson(*config)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %s", err)
	}

	return &Logger{loki: loki}, nil
}

func Debugf(format string, args ...interface{}) {
	pc, _, _, _ := runtime.Caller(1)
	callerName := runtime.FuncForPC(pc).Name()
	format = format + fmt.Sprintf(", callerName = (%s)\n", callerName)
	loggerInstance.loki.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	pc, _, _, _ := runtime.Caller(1)
	callerName := runtime.FuncForPC(pc).Name()
	format = format + fmt.Sprintf(", callerName = (%s)\n", callerName)
	loggerInstance.loki.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	pc, _, _, _ := runtime.Caller(1)
	callerName := runtime.FuncForPC(pc).Name()
	format = format + fmt.Sprintf(", callerName = (%s)\n", callerName)
	loggerInstance.loki.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	pc, _, _, _ := runtime.Caller(1)
	callerName := runtime.FuncForPC(pc).Name()
	format = format + fmt.Sprintf(", callerName = (%s)\n", callerName)
	loggerInstance.loki.Errorf(format, args...)
}

func Shutdown() {
	loggerInstance.loki.Shutdown()
}
